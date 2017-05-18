package main

import (
	"errors"
	"fmt"
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
	lastModifiedDate    string `json:"lastmodifieddate"`
	createDate   String `json:"createdate"`
}

type InsuranceInfo struct {
	provider      int `json:"provider"`
}

type ClaimInfo struct {
	claimInfo              string        `json:"id"`
	patientId             string        `json:"propertyId"`
	DateOfVisit           string        `json:"dateOfVisit"`
	NPI                   string        `json:"npi"`
	CPT                string        `json:"cpt"`
	ICD10                 string        `json:"icd10"`
	NDC                   string        `json:"ndc"`
	patientInfo           PatientInfo    `json:"personalInfo"`
	insuranceInfo          InsuranceInfo `json:"insuranceInfo"`
	Cost                 string        `json:"cost"`
	procedureStatus      String        `json:"procedureStatus"`
}

func GetLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering GetLoanApplication")

	if len(args) < 1 {
		logger.Error("Invalid number of arguments")
		return nil, errors.New("Missing loan application ID")
	}

	var loanApplicationId = args[0]
	bytes, err := stub.GetState(loanApplicationId)
	if err != nil {
		logger.Error("Could not fetch loan application with id "+loanApplicationId+" from ledger", err)
		return nil, err
	}
	return bytes, nil
}

func CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering CreateLoanApplication")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for loan application creation")
	}

	var loanApplicationId = args[0]
	var loanApplicationInput = args[1]

	err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))
	if err != nil {
		logger.Error("Could not save loan application to ledger", err)
		return nil, err
	}

	var customEvent = "{eventType: 'loanApplicationCreation', description:" + loanApplicationId + "' Successfully created'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully saved loan application")
	return nil, nil

}

/**
Updates the status of the loan application
**/
func UpdateLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	logger.Debug("Entering UpdateLoanApplication")

	if len(args) < 2 {
		logger.Error("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for loan application update")
	}

	var loanApplicationId = args[0]
	var status = args[1]

	laBytes, err := stub.GetState(loanApplicationId)
	if err != nil {
		logger.Error("Could not fetch loan application from ledger", err)
		return nil, err
	}
	var loanApplication LoanApplication
	err = json.Unmarshal(laBytes, &loanApplication)
	loanApplication.Status = status

	laBytes, err = json.Marshal(&loanApplication)
	if err != nil {
		logger.Error("Could not marshal loan application post update", err)
		return nil, err
	}

	err = stub.PutState(loanApplicationId, laBytes)
	if err != nil {
		logger.Error("Could not save loan application post update", err)
		return nil, err
	}

	var customEvent = "{eventType: 'loanApplicationUpdate', description:" + loanApplicationId + "' Successfully updated status'}"
	err = stub.SetEvent("evtSender", []byte(customEvent))
	if err != nil {
		return nil, err
	}
	logger.Info("Successfully updated loan application")
	return nil, nil

}

func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetLoanApplication" {
		return GetLoanApplication(stub, args)
	}
	return nil, nil
}

func GetCertAttribute(stub shim.ChaincodeStubInterface, attributeName string) (string, error) {
	logger.Debug("Entering GetCertAttribute")
	attr, err := stub.ReadCertAttribute(attributeName)
	if err != nil {
		return "", errors.New("Couldn't get attribute " + attributeName + ". Error: " + err.Error())
	}
	attrString := string(attr)
	return attrString, nil
}

func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateLoanApplication" {
		username, _ := GetCertAttribute(stub, "username")
		role, _ := GetCertAttribute(stub, "role")
		if role == "Bank_Home_Loan_Admin" {
			return CreateLoanApplication(stub, args)
		} else {
			return nil, errors.New(username + " with role " + role + " does not have access to create a loan application")
		}

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
