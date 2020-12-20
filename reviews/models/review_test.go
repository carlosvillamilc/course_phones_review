package models

import (
	"testing"
)

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars: stars,
		Comment: comment,
	}
}

func  Test_withCorrectsParams(t *testing.T)  {
	r := NewReview(4,"The ihpone X looks good")

	err := r.validate()

	if err != nil {
		t.Error("the validation did not pass")
		t.Fail()
	}	
}

func  Test_shouldFailWithWrongNumberOfStarts(t *testing.T)  {
	r := NewReview(10,"Excelent phone")

	err := r.validate()

	if err != nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}	
}