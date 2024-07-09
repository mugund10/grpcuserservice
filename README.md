## Assignment: Developing a Golang gRPC User Service with Search

> [!NOTE]
> This grpcUserService is hosted on aws and located @ `mugund10.openwaves.in:9292`
   * ANY gRPC lients can load this service methods/endpoints using server reflection
   * For more details see EXAMPLES from below

<details> 
<summary> INSTALLATION </summary>

### Docker


```shell
# Download image
docker pull mugundhan10/userservice:latest
# Run the service
docker run -d -p 9292:9292 mugundhan10/userservice:latest
```

### From source

```shell
# clone the repo
git clone https://github.com/mugund10/grpcuserservice.git

# build and run
cd grpcuserservice
go build  ./cmd/server/ && ./server

```


</details>



<details> 
<summary> EXAMPLES </summary>



## Using grpcurl

* ### To Install
    ```shell
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
    ```
* ### TO get Info
    *   About service
        ```shell
         #use localhost when running locally and the port will be same
         grpcurl --plaintext mugund10.openwaves.in:9292 list Userservice
        ```
         returns `
            Userservice.Fetchuser
            Userservice.Fetchusers
            Userservice.Search
            `

    *   About Endpoints
        ```shell
         grpcurl --plaintext mugund10.openwaves.in:9292 describe Userservice.Fetchuser
        ```
         returns `
            Userservice.Fetchuser is a method:
            rpc Fetchuser ( .Requestid ) returns ( .User );
            `
            
    *   About Requests
        ```shell
         grpcurl --plaintext mugund10.openwaves.in:9292 describe .Requestid
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
        # change your directory to cmd/client to access the json files
         grpcurl --plaintext -d @ mugund10.openwaves.in:9292 Userservice.Fetchuser < fetchuser.json
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
         grpcurl --plaintext -d @ mugund10.openwaves.in:9292 Userservice.Fetchusers < fetchusers.json
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
         grpcurl --plaintext -d @ mugund10.openwaves.in:9292 Userservice.Search < search3.json
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
            

</details>


<details> 
<summary> OBJECTIVES </summary>

1. Your task in this assessment is to create a Golang gRPC service that provides specific functionalities for managing user details and includes a search capability. The primary objectives are as follows:
2. Simulate a database by maintaining a list of user details within a variable.
3. Create gRPC endpoints for fetching user details based on a user ID and retrieving a list of 
   user details based on a list of user IDs.
4. Implement a search functionality to find user details based on specific criteria.

### Sample User Model:

{"id": 1, "fname": "Steve", "city": "LA", "phone": 1234567890, "height": 5.8, "Married": true}



### Maintain Code Quality and Design: 
1. Ensure that the code is well-structured and follows best practices.
2. Apply suitable design patterns to enhance the maintainability and extensibility of your service.



### Develop and Test APIs:
1. Implement the specified gRPC service methods to accomplish the tasks.
2. Write comprehensive unit tests to verify the correctness of your APIs.



### Input Validation and Response Handling:
1. Validate incoming requests to ensure they adhere to the defined requirements.
2. Handle requests appropriately and respond with meaningful messages, especially in the case of errors.



### Implement Search Functionality:
1. Create a search endpoint that allows users to search for specific user details based on criteria (e.g., city, phone number, marital status, etc.).



### Cover Edge Cases:
1. Identify potential edge cases and consider these in your implementation to provide robust and reliable functionality.



### Extra Brownie Points:
1. Use design patterns creatively to improve your service's architecture and efficiency.
2. Containerize the entire application using Docker for easy deployment and scaling.



### Mode of Submission:
Create a private GitHub repository for your project.
Include a README.md file in your repository with detailed instructions on how to:
   1. Build and run the application.
   2. Access the gRPC service endpoints.
   3. Provide any necessary configuration details.
 Share the link to your GitHub repository, and if necessary, provide a ZIP archive or a Google Drive link for submission.

 

Deadline: The deadline for submitting the assignment is in 4 days.

Location Requirement: This role is based in India, and we kindly request that only candidates located in India apply for this position.



Disclaimer: Please note that this assignment is solely for the purpose of assessment. Totality Corp will not use it for any other purpose than intended above.


</details>


- [x] completed all objectives
