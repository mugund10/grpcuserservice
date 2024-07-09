## Using grpcurl

* ### To Install
    ```shell
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
    ```
* ### TO get Info
    *   About service
        ```shell
         grpcurl --plaintext localhost:9292 list Userservice
        ```
         returns `
            Userservice.Fetchuser
            Userservice.Fetchusers
            Userservice.Search
            `

    *   About Endpoints
        ```shell
         grpcurl --plaintext localhost:9292 describe Userservice.Fetchuser
        ```
         returns `
            Userservice.Fetchuser is a method:
            rpc Fetchuser ( .Requestid ) returns ( .User );
            `
            
    *   About Requests
        ```shell
         grpcurl --plaintext localhost:9292 describe .Requestid
        ```
         returns `
            Requestid is a message:
            message Requestid {
             int64 id = 1;
            }
            `
* ### TO Call a gRPCServer

    *   Endpoint Fetchuser
        ```shell
         grpcurl --plaintext -d @ localhost:9292 Userservice.Fetchuser < fetchuser.json
        ```
         returns `
            {
                 "id": "10",
                 "fname": "Axar Patel",
                 "city": "AHMEDABAD",
                 "phone": "9493423245",
                 "height": 6,
                 "married": true
            }
            `
        
        where Fetchuser.json contains 
        ``  {
            "id": 10
            }
        ``

    *   Endpoint Fetchusers
        ```shell
         grpcurl --plaintext -d @ localhost:9292 Userservice.Fetchusers < fetchusers.json
        ```
         returns `
                        {
            "user": [
                {
                "id": "6",
                "fname": "Jasprit Bumrah",
                "city": "AHMEDABAD",
                "phone": "9467633467",
                "height": 6.6
                },
                {
                "id": "4",
                "fname": "Rishabh Pant",
                "city": "ROORKEE",
                "phone": "9323891212",
                "height": 5.8
                },
                {
                "id": "3",
                "fname": "Virat Kohli",
                "city": "DELHI",
                "phone": "9275843489",
                "height": 6.4,
                "married": true
                },
                {
                "id": "2",
                "fname": "Yashasvi Jaiswal",
                "city": "BHADOHI",
                "phone": "9014313478",
                "height": 5.3
                }
            ]
            }
            `
        
        where Fetchusers.json contains 
        ``  {
            "id": [
                {
                    "id": "6"
                },
                {
                    "id": "4"
                },
                {
                    "id": "3"
                },
                {
                    "id": "2"
                }
            ]
           }
        ``

    *   Endpoint Search
        ```shell
         grpcurl --plaintext -d @ localhost:9292 Userservice.Search < search3.json
        ```
         returns `
                        {
            "user": [
                {
                "id": "1",
                "fname": "Rohit Sharma",
                "city": "MUMBAI",
                "phone": "903248732",
                "height": 5.2,
                "married": true
                },
                {
                "id": "3",
                "fname": "Virat Kohli",
                "city": "DELHI",
                "phone": "9275843489",
                "height": 6.4,
                "married": true
                },
                {
                "id": "7",
                "fname": "Ravindra Jadeja",
                "city": "AHMEDABAD",
                "phone": "9219028912",
                "height": 5.3,
                "married": true
                },
                {
                "id": "10",
                "fname": "Axar Patel",
                "city": "AHMEDABAD",
                "phone": "9493423245",
                "height": 6,
                "married": true
                },
                {
                "id": "11",
                "fname": "Hrithik Roshan",
                "city": "MUMBAI",
                "phone": "9032433732",
                "height": 5.2,
                "married": true
                },
                {
                "id": "13",
                "fname": "ShahRukh Khan",
                "city": "DELHI",
                "phone": "9278448912",
                "height": 6.4,
                "married": true
                },
                {
                "id": "14",
                "fname": "Ajay Devgan",
                "city": "DELHI",
                "phone": "9322321278",
                "height": 5.8,
                "married": true
                },
                {
                "id": "15",
                "fname": "Kapil Dev",
                "city": "FAZILKA",
                "phone": "9034622342",
                "height": 5.1,
                "married": true
                },
                {
                "id": "16",
                "fname": "Navdeep Asija",
                "city": "FAZILKA",
                "phone": "9464563467",
                "height": 6.6,
                "married": true
                },
                {
                "id": "18",
                "fname": "Ram Nath Kovind",
                "city": "KANPUR",
                "phone": "9356324460",
                "height": 5.9,
                "married": true
                },
                {
                "id": "19",
                "fname": "Sundar Pichai",
                "city": "CHENNAI",
                "phone": "9023435832",
                "height": 6.2,
                "married": true
                },
                {
                "id": "20",
                "fname": "Lady Andal amma",
                "city": "CHENNAI",
                "phone": "9432423245",
                "height": 6,
                "married": true
                }
            ]
            }
            `
        
        where Fetchuser.json contains 
        `` {
            "field": "married",
            "value": {
            "@type": "type.googleapis.com/google.protobuf.BoolValue",
            "value": true
            }
          }``
            