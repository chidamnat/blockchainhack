package main

import (
    "errors"
    "fmt"
    "strconv"
	  "encoding/json"
    "strings"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("mylogger")

type SampleChaincode struct {
}

//custom data models
type PatientInfo struct {
    PatientId string `json:"patientid"`
    Zipcd  string `json:"zipcd"`
    State       string `json:"state"`
    Birthdate     string `json:"birthdate"`
    LastModifiedDate    string `json:"lastmodifieddate"`
    CreateDate   string `json:"createdate"`
}

type ClaimInfo struct {
    ClaimInfo             string       `json:"id"`
    PatientId             string       `json:"patientId"`
    DateOfVisit           string       `json:"dateOfVisit"`
    NPI                   string       `json:"npi"`
    CPT                   string       `json:"cpt"`
    ICD10                 string       `json:"icd10"`
    NDC                   string       `json:"ndc"`
    //PatientInfo           PatientInfo    `json:"personalInfo"`
    PatientInfo           string       `json:"personalInfo"`
    //InsuranceInfo          InsuranceInfo `json:"insuranceInfo"`
    InsuranceInfo        string        `json:"insuranceInfo"`
    Cost                 string        `json:"cost"`
    ProcedureStatus      string        `json:"procedureStatus"`
}
var accountIndexStr = "_accountindex"	  // Define an index varibale to track all the accounts stored in the world state

func GetPatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering GetPatientInfo")

    if len(args) < 1 {
        logger.Error("Invalid number of arguments")
        return nil, errors.New("Missing patient ID")
    }

    var patientID = args[0]
    bytes, err := stub.GetState(patientID)
    if err != nil {
        logger.Error("Could not fetch patient info with id "+patientID+" from ledger", err)
        return nil, err
    }
    return bytes, nil
}

func GetCompletePatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering GetCompletePatientInfo")

    if len(args) < 1 {
        logger.Error("Invalid number of arguments")
        return nil, errors.New("Missing patient ID")
    }

    var patientID = args[0]
    // bytes, err := stub.GetState(patientID)
    // if err != nil {
    //     logger.Error("Could not fetch patient info with id "+patientID+" from ledger", err)
    //     return nil, err
    // }
    // return bytes, nil
    //get the account index
  	accountsAsBytes, err := stub.GetState(accountIndexStr)
  	if err != nil {
  		return nil, errors.New("Failed to get account index")
  	}
  	var accountIndex []string
  	json.Unmarshal(accountsAsBytes, &accountIndex)

    for i,val := range accountIndex{
        var m ClaimInfo
        claimOut, _ = GetClaimInfo(stub,strings.Fields(val))
        json.Unmarshal(claimOut, &m)
        //json.Unmarshal(GetClaimInfo(stub,val), &m)
		    if m.patientId == patientId {
          accountIndex = append(accountIndex[i])
        }
    }
    jsonAsBytes, _ := json.Marshal(accountIndex)
    return jsonAsBytes, nil
}

func GetClaimInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering GetClaimInfo")

    if len(args) < 1 {
        logger.Error("Invalid number of arguments")
        return nil, errors.New("Missing Claim ID")
    }

    var claimId = args[0]
    bytes, err := stub.GetState(claimId)
    if err != nil {
        logger.Error("Could not fetch patient info with id "+claimId+" from ledger", err)
        return nil, err
    }
    return bytes, nil
}

func CreatePatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering CreatePatientInfo")

    if len(args) < 6 {
        logger.Error("Invalid number of args")
        return nil, errors.New("Expected at least two arguments for patient info creation.")
    }

    var patientID = args[0]
    //var patientInfoInput = args[1]
    var zipcd = args[1]
    var state = args[2]
    var birthdate = args[3]
    var lastmodifieddate = args[4]
    var createdate = args[5]
    patientInfoInput := `{ "zipcd": "` + zipcd + `","state": "` + state + `", "birthdate": "` + birthdate + `", "lastmodifieddate": "` + lastmodifieddate + `" ,"createdate": "` + createdate + `"}`

    err := stub.PutState(patientID, []byte(patientInfoInput))
    if err != nil {
        logger.Error("Could not save patient info to ledger", err)
        return nil, err
    }
    logger.Info("Successfully saved patient info.")
    return nil, nil
}

func CreateClaimInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering CreateClaimInfo")
    if len(args) < 10 {
        logger.Error("Invalid number of args")
        return nil, errors.New("Expected at least two arguments for claim info creation.")
    }
    var claimId = args[0]
    var patientId = args[1]
    var dateOfVisit = args[2]
    var npi = args[3]
    var cpt = args[4]
    var icd10 = args[5]
    var ndc = args[6]
    //var personalInfo = GetPatientInfo(stub, args[7])
    var insuranceInfo = args[7]
    var cost = args[8]
    var procedureStatus = args[9]

    claimInfoInput := `{ "patientId": "` + patientId + `", "dateOfVisit": "` + dateOfVisit + `", "npi": "` + npi + `", "cpt": "` + cpt + `", "icd10": "` + icd10 + `", "ndc": "` + ndc +  `","insuranceInfo": "` + insuranceInfo + `", "cost": "` + cost + `", "procedureStatus": "` + procedureStatus + `"}`
    //var claimInfoInput = args[1]

    err := stub.PutState(claimId, []byte(claimInfoInput))
    if err != nil {
        logger.Error("Could not save claim info to ledger", err)
        return nil, err
    }

    //get the account index
  	accountsAsBytes, err := stub.GetState(accountIndexStr)
  	if err != nil {
  		return nil, errors.New("Failed to get account index")
  	}
  	var accountIndex []string
  	json.Unmarshal(accountsAsBytes, &accountIndex)

    //append the index
  	accountIndex = append(accountIndex, claimId)
  	jsonAsBytes, _ := json.Marshal(accountIndex)
  	err = stub.PutState(accountIndexStr, jsonAsBytes)

    logger.Info("Successfully saved patient info.")
    return nil, nil

}

func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    var Aval int
    var err error

    if len(args) != 1 {
      return nil, errors.New("Incorrect number of arguments. Expecting a single integer")
    }

    // Initialize the chaincode
    Aval, err = strconv.Atoi(args[0])
    if err != nil {
      return nil, errors.New("Expecting integer value for testing the blockchain network")
    }

    // Write the state to the ledger, test the network
    err = stub.PutState("test_key", []byte(strconv.Itoa(Aval)))
    if err != nil {
      return nil, err
    }

    var empty []string
    jsonAsBytes, _ := json.Marshal(empty)								//marshal an emtpy array of strings to clear the account index
    err = stub.PutState(accountIndexStr, jsonAsBytes)
    if err != nil {
      return nil, err
    }

    return nil, nil
}


func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if function == "GetPatientInfo" {
        return GetPatientInfo(stub, args)
    }
    if function == "GetClaimInfo" {
        return GetClaimInfo(stub, args)
    }
    return nil, nil
}


func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if function == "CreatePatientInfo" {
        return CreatePatientInfo(stub, args)
    }

    if function == "CreateClaimInfo" {
        return CreateClaimInfo(stub, args)
    }
    return nil, nil
}


func main() {

    lld, _ := shim.LogLevel("DEBUG")
    fmt.Println(lld)

    logger.SetLevel(lld)
    fmt.Println(logger.IsEnabledFor(lld))

    err := shim.Start(new(SampleChaincode))
    if err != nil {
        logger.Error("Could not start SampleChaincode")
    } else {
        logger.Info("SampleChaincode successfully started")
    }

}
