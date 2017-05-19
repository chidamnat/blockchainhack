package main

import (
    "time"
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
    PatientId             string       `json:"patientid"`
    Zipcd                 string       `json:"zipcd"`
    State                 string       `json:"state"`
    Birthdate             string       `json:"birthdate"`
    ClaimHistory          []string     `json:claimhistory`
    LastModifiedDate      time.Time    `json:"lastmodifieddate"`
    CreatedDate           time.Time    `json:"createdate"`
}

type ClaimInfo struct {
    ClaimInfo             string       `json:"id"`
    PatientId             string       `json:"patientId"`
    DateOfVisit           string       `json:"dateOfVisit"`
    NPI                   string       `json:"npi"`
    CPT                   string       `json:"cpt"`
    ICD10                 string       `json:"icd10"`
    NDC                   string       `json:"ndc"`
    PatientInfo           string       `json:"personalInfo"`
    InsuranceInfo         string        `json:"insuranceInfo"`
    Cost                  string        `json:"cost"`
    ProcedureStatus       string        `json:"procedureStatus"`
}

func GetPatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering GetPatientInfo")

    if len(args) < 1 {
        logger.Error("Invalid number of arguments")
        return nil, errors.New("Missing patient ID")
    }

    var patientID = args[0]
    patientInfoJSON, err := stub.GetState(patientID)
    if err != nil {
        logger.Error("Could not fetch patient info with id "+patientID+" from ledger", err)
        return nil, err
    }
    return patientInfoJSON, nil
}

func GetPatientClaimHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering GetCompletePatientInfo")

    if len(args) < 1 {
        logger.Error("Invalid number of arguments")
        return nil, errors.New("Missing patient ID")
    }

    patientID := args[0]
    var patientInfo PatientInfo

    patientInfoRaw, err := stub.GetState(patientID)
    json.Unmarshal(patientInfoRaw, &patientInfo)
    if err != nil {
        logger.Error("Could not fetch patient info with id "+patientID+" from ledger", err)
        return nil, err
    }

    var claimHistoryArr []string
    var claimID string
    var claimInfoRaw []byte
    var claimInfoJSON string

    for i := 0; i < len(patientInfo.ClaimHistory); i++ {
          claimID = patientInfo.ClaimHistory[i]
          claimInfoRaw, err = stub.GetState(claimID)
          claimInfoJSON = string(claimInfoRaw)
          if err != nil {
            logger.Error("Could not fetch claim info with id "+claimID+" from ledger", err)
            return nil, err
          }
          claimHistoryArr = append(claimHistoryArr, claimInfoJSON)
    }

    var claims []byte
    claims, err = json.Marshal(claimHistoryArr)
    return claims, nil
}

func GetClaimInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering GetClaimInfo")

    if len(args) < 1 {
        logger.Error("Invalid number of arguments")
        return nil, errors.New("Missing Claim ID")
    }

    var claimId = args[0]
    claimInfoJSON, err := stub.GetState(claimId)
    if err != nil {
        logger.Error("Could not fetch patient info with id "+claimId+" from ledger", err)
        return nil, err
    }

    return claimInfoJSON, nil
}

func CreatePatientInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    logger.Debug("Entering CreatePatientInfo")

    if len(args) < 6 {
        logger.Error("Invalid number of args")
        return nil, errors.New("Expected at least two arguments for patient info creation.")
    }
    var patientID = args[0]

    var patientInfo *PatientInfo

    patientInfo = &PatientInfo{
      PatientId: patientID,
      Zipcd: args[1],
      State: args[2],
      Birthdate: args[3],
      ClaimHistory: make([]string, 0),
      LastModifiedDate: time.Now(),
      CreatedDate: time.Now(),
    }

    patientInfoJSON, err := json.Marshal(patientInfo)

    err = stub.PutState(patientID, patientInfoJSON)
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

    var claimID = args[0]
    var patientID = args[1]

    var claimInfo = &ClaimInfo{
      ClaimInfo: claimID,
      PatientId: patientID,
      DateOfVisit: args[2],
      NPI: args[3],
      CPT: args[4],
      ICD10: args[5],
      NDC: args[6],
      InsuranceInfo: args[7],
      Cost: args[8],
      ProcedureStatus: args[9],
    }

    claimInfoJSON, err := json.Marshal(claimInfo)
    err = stub.PutState(claimID, claimInfoJSON)
    if err != nil {
        logger.Error("Could not save claim info to ledger", err)
        return nil, err
    }

    //Update PatientInfo.ClaimHistory
    patientInfoJSONOld, err := stub.GetState(patientID)
    if err != nil {
        logger.Error("Could not retrieve patient info", err)
        return nil, err
    }

    var patientInfo PatientInfo
    json.Unmarshal(patientInfoJSONOld, &patientInfo)
    patientInfo.ClaimHistory = append(patientInfo.ClaimHistory, claimID)
    patientInfo.LastModifiedDate = time.Now()
    patientInfoJSONNew, err := json.Marshal(patientInfo)
    err = stub.PutState(patientID, patientInfoJSONNew)
    if err != nil {
        logger.Error("Could not save patient info to ledger", err)
        return nil, err
    }
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

    return nil, nil
}


func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if function == "GetPatientInfo" {
        return GetPatientInfo(stub, args)
    }
    if function == "GetClaimInfo" {
        return GetClaimInfo(stub, args)
    }
    if function == "GetPatientClaimHistory" {
        return GetPatientClaimHistory(stub, args)
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
