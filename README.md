'''
    package main
    
    import (
            "fmt"
            "github.com/dimanlin/salesforce_martiix_client"
    )
    
    func main() {
            client := salesforce_martiix_client.CreateMClient("HOST", "EMAIL", "PASSWORD", "APIVersion")
            fmt.Println("TEST GetAllMarttix")
            allMartix := client.GetAllMarttix()
    
            fmt.Println("TEST Marttix Id exist")
            fmt.Println((*allMartix)[0].IdExist())
    
            allClient, _ := client.GetAllClients()
            fmt.Println("TEST GetClientAll")
            for _, record := range *allClient {
                    fmt.Println(record.ContactEmail)
            }
    
            fmt.Println("TEST GetClientByEmail")
            userClient, _ := client.GetClientByEmail("uitest@mail.com")
            fmt.Println(userClient)
    }
'''