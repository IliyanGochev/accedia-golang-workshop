package main

import (  
    "fmt"
)


type employee struct {  
    salary  int
    country string
}

func main() {
    emp1 := employee{
        salary: 1000,
        country: "USA",
	}

	employeeInfo := map[string]employee{
        "Steve": emp1,
    }

    emp2 := employee{
        salary: 2000,
        country: "France",
    }

    employeeInfo["Jean"] = emp2
    val, ok := employeeInfo["Jean"]
    fmt.Println("val, ok:", val, ok) // val, ok: {2000 France} true

    fmt.Println("Jean", employeeInfo["Jean"]) // Jean {2000 France}
    fmt.Println("Steve", employeeInfo["Steve"]) // Steve {1000 USA}


    delete(employeeInfo, "Steve")
    fmt.Println("Steve", employeeInfo["Steve"]) // Steve {0 }
}