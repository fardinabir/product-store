# Product Store

## Project Overview

This project consists of a set of APIs built to meet specific requirements handling a product store. Below are some key points and instructions related to the assignment.

### Task Feedback

1. **API Implementation:** All APIs have been implemented according to the specified requirements.
2. **Pagination API:** Access the pagination API at [http://localhost:8085/api/products?min_price=700&max_price=2000&product_name=Grey Fatua](http://localhost:8085/api/products?min_price=700&max_price=2000&product_name=Grey%20Fatua).
3. **Tree Generator Endpoint:** Find the tree generator endpoint at [http://localhost:8085/api/categories/get-tree](http://localhost:8085/api/categories/get-tree).
4. **LRU Cache Problem Solution:** The solution to the LRU cache problem can be found in the `problem_solving` directory.

### Project Details

- **API Documentation:** Swagger documentation has been maintained. The documentation files are located in the `/docs` directory.
- **Commands for API Details:**
    1. Run `swagger generate spec -w . -o docs/swagger.json` to generate the Swagger specification.
    2. Run `swagger serve --port 8090 -F=swagger docs/swagger.json` to serve the Swagger documentation. Access it at [http://localhost:8090/docs](http://localhost:8090/docs).

**Note:** Due to time constraints, couldn't prepare appropriate Dockerfile with a swagger container, some functionalities also left incomplete. Postman-collection is also shared in root directory.

