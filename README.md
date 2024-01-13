# Monday API Tools
This code implements parsers for the following:
1. Webhook message parser
2. Struct ORM for Monday's GraphQL API

## Usecase
If you wish to setup an auxilery system which recieves webhook from Monday updates and then map it to an external DB you can use this tool.

For GraphQL responses, just create a struct which describes your Monday table schema and it will automagially bind to the response json.

### Example Diagram

T