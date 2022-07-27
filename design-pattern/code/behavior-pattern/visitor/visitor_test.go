package visitor

import "testing"

func TestRequestVisitor(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServiceRequestVisitor{})
}

func TestAnalysis(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Accept(&AnalysisVisitor{})
}
