package main

import (
    "errors"
    "fmt"
    "strconv"
	  "encoding/json"
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

type InsuranceInfo struct {
    Provider      int `json:"provider"`
}

type ClaimInfo struct {
    ClaimInfo              string        `json:"id"`
    PatientId             string        `json:"propertyId"`
    DateOfVisit           string        `json:"dateOfVisit"`
    NPI                   string        `json:"npi"`
    CPT                string        `json:"cpt"`
    ICD10                 string        `json:"icd10"`
    NDC                   string        `json:"ndc"`
    PatientInfo           PatientInfo    `json:"personalInfo"`
    InsuranceInfo          InsuranceInfo `json:"insuranceInfo"`
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

func CreatePatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering CreatePatientInfo")

    if len(args) < 2 {
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
    patientInfoInput := `{ "state": "` + state + `", "birthdate": "` + birthdate + `", "lastmodifieddate": "` + lastmodifieddate + `" }`


    err := stub.PutState(patientID, []byte(patientInfoInput))
    if err != nil {
        logger.Error("Could not save patient info to ledger", err)
        return nil, err
    }

    // var customEvent = "{eventType: 'patientInfoCreation', description: " + patientID + "' Successfully created'}"
    // err = stub.SetEvent("evtSender", []byte(customEvent))
    // if err != nil {
    //     return nil, err
    // }
    logger.Info("Successfully saved patient info.")
    return nil, nil

}

func CreateClaimInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering CreateClaimInfo")

    if len(args) < 2 {
        logger.Error("Invalid number of args")
        return nil, errors.New("Expected at least two arguments for claim info creation.")
    }

    var claimId = args[0]
    var claimInfoInput = args[1]

    err := stub.PutState(claimId, []byte(claimInfoInput))
    if err != nil {
        logger.Error("Could not save claim info to ledger", err)
        return nil, err
    }

    // var customEvent = "{eventType: 'claimInfoCreation', description: " + claimId + "' Successfully created'}"
    // err = stub.SetEvent("evtSender", []byte(customEvent))
    // if err != nil {
    //     return nil, err
    // }
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
    return nil, nil
}

// func GetCertAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
//     logger.Debug("Entering GetCertAttribute")
//     attr, err := stub.ReadCertAttribute(attributeName)
//     if err != nil {
//         return "", errors.New("Couldn't get attribute " + attributeName + ". Error: " + err.Error())
//     }
//     attrString := string(attr)
//     return attrString, nil
// }

func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if function == "CreatePatientInfo" {
        return CreatePatientInfo(stub, args)
    }

    if function == "CreateClaimInfo" {
        return CreateClaimInfo(stub, args)
    }
    return nil, nil
}

type customEvent struct {
    Type       string `json:"type"`
    Decription string `json:"description"`
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
