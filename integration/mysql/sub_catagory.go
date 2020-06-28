package mysql

type SubCategoryEntity struct {
	SC_Id          string
	SC_Name        string
	SC_Description string
	SC_ImageUrl    string
	SC_IsActive    bool
	BC_Id          string
}

// Functions to Read & Update the SubCategoryEntity
