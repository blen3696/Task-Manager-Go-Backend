| Property | Value |
| --- | --- |
| event | _complex array_ |
| info | _object_ |
| item | _complex array_ |

### event
| # | listen| script|
| --- | --- | --- |
| 1 | test | _complex_ |

#### event[0]
| Property | Value |
| --- | --- |
| listen | test |
| script | _object_ |

##### event[0].script
| Property | Value |
| --- | --- |
| exec | [// ============================================, // COLLECTION-LEVEL TESTS - Applied to ALL requests, // ============================================, , // Health Check 1: Response Time, pm.test('Response time is acceptable (< 5000ms)', function () {,     pm.expect(pm.response.responseTime).to.be.below(5000);, });, , // Health Check 2: Valid JSON Response, pm.test('Response is valid JSON', function () {,     try {,         pm.response.json();,     } catch (e) {,         pm.expect.fail('Response is not valid JSON: ' + e.message);,     }, });, , // Health Check 3: Content-Type Header, pm.test('Content-Type header is present', function () {,     pm.response.to.have.header('Content-Type');, });, , // Health Check 4: No Server Errors, pm.test('No server errors (5xx)', function () {,     pm.expect(pm.response.code).to.be.below(500);, });, , // Health Check 5: Authentication Check (if token is used), pm.test('Request is authenticated (no 401/403)', function () {,     pm.expect(pm.response.code).to.not.be.oneOf([401, 403]);, });, , // Workflow Variable Logging (for debugging), if (pm.response.code >= 200 && pm.response.code < 300) {,     console.log('Request: ' + pm.info.requestName + ' - Status: ' + pm.response.code);, }] |
| type | text/javascript |

### info
| Property | Value |
| --- | --- |
| _exporter_id | 40728853 |
| _postman_id | 37baae72-bd01-4c95-b652-8194f8483504 |
| description | ## TASK SERVICE API<br><br>### Overview<br><br>This collection contains requests for the **TASK SERVICE API**. It is intended for internal/team use when developing and testing task-related functionality.<br><br>---<br><br>### Environment Variables Used<br><br>From **Local Development** environment:<br><br>- `base_url`<br>    <br>    - Example usage: `{{base_url}}/tasks`<br>        <br>    - Purpose: Base URL for all TASK SERVICE API endpoints.<br>        <br>- `token`<br>    <br>    - Example usage (if applicable): `Authorization: Bearer {{token}}`<br>        <br>    - Purpose: Holds the auth token used to access secured endpoints of the API.<br>        <br><br>---<br><br>### Notes for the Team<br><br>- As additional endpoints are added to this collection, update this description with:<br>    - New requests and their purpose<br>        <br>    - Required headers and authentication scheme<br>        <br>    - Request body schemas and examples<br>        <br>    - At least one example successful response per request.<br>        <br>- When the real URLs, methods, and structures are known, replace the assumed values above (especially for **Create Task**) with the actual API definitions. |
| name | TASK SERVICE API |
| schema | https://schema.getpostman.com/json/collection/v2.1.0/collection.json |

### item
| # | event| name| request| response|
| --- | --- | --- | --- | --- |
| 1 | _complex_ | Create Task | _complex_ | _complex_ |
| 2 | _complex_ | Get Tasks | _complex_ | _complex_ |
| 3 | _complex_ | Get Task | _complex_ | _complex_ |
| 4 |  | Delete Task | _complex_ | _complex_ |
| 5 | _complex_ | Update Task | _complex_ | _complex_ |

#### item[0]
| Property | Value |
| --- | --- |
| event | _complex array_ |
| name | Create Task |
| request | _object_ |
| response | _complex array_ |

##### item[0].event
| # | listen| script|
| --- | --- | --- |
| 1 | test | _complex_ |

###### item[0].event[0]
| Property | Value |
| --- | --- |
| listen | test |
| script | _object_ |

####### item[0].event[0].script
| Property | Value |
| --- | --- |
| exec | [pm.test('Create Task - Response has task ID', function () {,     const jsonData = pm.response.json();,     pm.expect(jsonData).to.have.property('id');,     if (jsonData.id) {,         pm.collectionVariables.set('created_task_id', jsonData.id);,     }, }), , pm.test('Create Task - Response has required task fields', function () {,     const jsonData = pm.response.json();,     pm.expect(jsonData).to.have.property('title');, }), , pm.test('Create Task - Status code is 201 Created or 200 OK', function () {,     pm.expect(pm.response.code).to.be.oneOf([,         200,,         201,     ]);, })] |
| packages | _object_ |
| requests | _object_ |
| type | text/javascript |

######## item[0].event[0].script.packages

_Empty object_

######## item[0].event[0].script.requests

_Empty object_

##### item[0].request
| Property | Value |
| --- | --- |
| body | _object_ |
| description | ### Create Task<br><br>Creates a new task resource in the Task Service.<br><br>**Request URL**<br>- `POST {{base_url}}/tasks`<br><br>**Request Body (JSON)**<br><br>```json<br>{<br>  "title": "Finish DSA",<br>  "description": "Finish on the BFS and DFS",<br>  "due_date": "2025-01-20T00:00:00Z",<br>  "status": "pending"<br>}<br>```<br><br>**Fields**<br>- `title` (string, required)<br>  - Short name of the task.<br>- `description` (string, optional/recommended)<br>  - More detailed information about what needs to be done.<br>- `due_date` (string, required, ISO 8601 datetime)<br>  - Due date/time for the task, for example: `2025-01-20T00:00:00Z`.<br>- `status` (string, required)<br>  - Current state of the task.<br>  - Typical values include: `"pending"`, `"completed"` (and any other valid states supported by the API).<br><br>**Successful Responses**<br>- Status codes: **200 OK** or **201 Created**<br>- The response body includes:<br>  - A generated `id` for the newly created task.<br>  - Echoed core task fields such as `title`, and may include `description`, `due_date`, `status`, and other task metadata depending on the implementation. |
| header | [] |
| method | POST |
| url | _object_ |

###### item[0].request.body
| Property | Value |
| --- | --- |
| mode | raw |
| options | _object_ |
| raw | {<br>  "title": "Finish DSA",<br>  "description": "Finish on the BFS and DFS",<br>  "due_date": "2025-01-20T00:00:00Z",<br>  "status": "pending"<br>}<br><br> |

####### item[0].request.body.options
| Property | Value |
| --- | --- |
| raw | _object_ |

######## item[0].request.body.options.raw
| Property | Value |
| --- | --- |
| language | json |

###### item[0].request.url
| Property | Value |
| --- | --- |
| host | [{{base_url}}] |
| path | [tasks] |
| raw | {{base_url}}/tasks |

##### item[0].response
| # | _postman_previewlanguage| body| code| cookie| header| name| originalRequest| status|
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | null | {<br>    "id": 0,<br>    "title": "Finish DSA",<br>    "description": "Finish on the BFS and DFS",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "pending"<br>} | 201 | [] | _complex_ | Create Task | _complex_ | Created |

###### item[0].response[0]
| Property | Value |
| --- | --- |
| _postman_previewlanguage | null |
| body | {<br>    "id": 0,<br>    "title": "Finish DSA",<br>    "description": "Finish on the BFS and DFS",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "pending"<br>} |
| code | 201 |
| cookie | [] |
| header | _complex array_ |
| name | Create Task |
| originalRequest | _object_ |
| status | Created |

####### item[0].response[0].header
| # | key| value|
| --- | --- | --- |
| 1 | Content-Type | application/json; charset=utf-8 |
| 2 | Date | Thu, 18 Dec 2025 10:59:54 GMT |
| 3 | Content-Length | 155 |

####### item[0].response[0].originalRequest
| Property | Value |
| --- | --- |
| body | _object_ |
| header | [] |
| method | POST |
| url | _object_ |

######## item[0].response[0].originalRequest.body
| Property | Value |
| --- | --- |
| mode | raw |
| options | _object_ |
| raw | {<br>  "title": "Finish DSA",<br>  "description": "Finish on the BFS and DFS",<br>  "due_date": "2025-01-20T00:00:00Z",<br>  "status": "pending"<br>}<br><br> |

######### item[0].response[0].originalRequest.body.options
| Property | Value |
| --- | --- |
| raw | _object_ |

########## item[0].response[0].originalRequest.body.options.raw
| Property | Value |
| --- | --- |
| language | json |

######## item[0].response[0].originalRequest.url
| Property | Value |
| --- | --- |
| path | [tasks] |
| raw | /tasks |

#### item[1]
| Property | Value |
| --- | --- |
| event | _complex array_ |
| name | Get Tasks |
| request | _object_ |
| response | _complex array_ |

##### item[1].event
| # | listen| script|
| --- | --- | --- |
| 1 | test | _complex_ |

###### item[1].event[0]
| Property | Value |
| --- | --- |
| listen | test |
| script | _object_ |

####### item[1].event[0].script
| Property | Value |
| --- | --- |
| exec | [pm.test('Get Tasks - Response is an array', function () {,     const jsonData = pm.response.json();,     const tasks = Array.isArray(jsonData) ? jsonData : jsonData.tasks \|\| jsonData.data \|\| [];,     pm.expect(tasks).to.be.an('array');, }), , pm.test('Get Tasks - Each task has required fields', function () {,     const jsonData = pm.response.json();,     const tasks = Array.isArray(jsonData) ? jsonData : jsonData.tasks \|\| jsonData.data \|\| [];,     if (tasks.length > 0) {,         tasks.forEach(function (task, index) {,             pm.expect(task, 'Task at index ' + index + ' should have id').to.have.property('id');,             pm.expect(task, 'Task at index ' + index + ' should have title').to.have.property('title');,         });,     }, }), , pm.test('Get Tasks - Status code is 200 OK', function () {,     pm.response.to.have.status(200);, })] |
| packages | _object_ |
| requests | _object_ |
| type | text/javascript |

######## item[1].event[0].script.packages

_Empty object_

######## item[1].event[0].script.requests

_Empty object_

##### item[1].request
| Property | Value |
| --- | --- |
| description | ### Get Tasks<br><br>Retrieve a list of tasks for the current context.<br><br>**Method & URL**  <br>`GET {{base_url}}/tasks`<br><br>**Response Format**  <br>Returns an array of task objects. Each task includes at least:<br>- `id` (unique identifier)<br>- `title`<br><br>Tasks may also include additional fields such as:<br>- `description`<br>- `due_date`<br>- `status`<br><br>**Tests**  <br>Current tests validate that:<br>- The response status code is `200 OK`.<br>- The response body is an array of tasks.<br>- Each task in the array contains `id` and `title` properties. |
| header | [] |
| method | GET |
| url | _object_ |

###### item[1].request.url
| Property | Value |
| --- | --- |
| host | [{{base_url}}] |
| path | [tasks] |
| raw | {{base_url}}/tasks |

##### item[1].response
| # | _postman_previewlanguage| body| code| cookie| header| name| originalRequest| status|
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | null | [<br>    {<br>        "id": 1,<br>        "title": "Task 1",<br>        "description": "Description of Task 1",<br>        "due_date": "2025-12-18T13:21:15.3573694+03:00",<br>        "status": "Pending"<br>    },<br>    {<br>        "id": 2,<br>        "title": "Task 2",<br>        "description": "Description of Task 2",<br>        "due_date": "2025-12-19T13:21:15.3573694+03:00",<br>        "status": "Completed"<br>    },<br>    {<br>        "id": 3,<br>        "title": "Task 3",<br>        "description": "Description of Task 3",<br>        "due_date": "2025-12-20T13:21:15.3578734+03:00",<br>        "status": "In Progress"<br>    },<br>    {<br>        "id": 0,<br>        "title": "Finish report",<br>        "description": "Prepare Q1 report",<br>        "due_date": "2025-01-20T00:00:00Z",<br>        "status": "completed"<br>    },<br>    {<br>        "id": 0,<br>        "title": "Finish report",<br>        "description": "Prepare Q1 report",<br>        "due_date": "2025-01-20T00:00:00Z",<br>        "status": "pending"<br>    },<br>    {<br>        "id": 0,<br>        "title": "Finish DSA",<br>        "description": "Finish on the BFS and DFS",<br>        "due_date": "2025-01-20T00:00:00Z",<br>        "status": "pending"<br>    }<br>] | 200 | [] | _complex_ | Get Tasks | _complex_ | OK |

###### item[1].response[0]
| Property | Value |
| --- | --- |
| _postman_previewlanguage | null |
| body | [<br>    {<br>        "id": 1,<br>        "title": "Task 1",<br>        "description": "Description of Task 1",<br>        "due_date": "2025-12-18T13:21:15.3573694+03:00",<br>        "status": "Pending"<br>    },<br>    {<br>        "id": 2,<br>        "title": "Task 2",<br>        "description": "Description of Task 2",<br>        "due_date": "2025-12-19T13:21:15.3573694+03:00",<br>        "status": "Completed"<br>    },<br>    {<br>        "id": 3,<br>        "title": "Task 3",<br>        "description": "Description of Task 3",<br>        "due_date": "2025-12-20T13:21:15.3578734+03:00",<br>        "status": "In Progress"<br>    },<br>    {<br>        "id": 0,<br>        "title": "Finish report",<br>        "description": "Prepare Q1 report",<br>        "due_date": "2025-01-20T00:00:00Z",<br>        "status": "completed"<br>    },<br>    {<br>        "id": 0,<br>        "title": "Finish report",<br>        "description": "Prepare Q1 report",<br>        "due_date": "2025-01-20T00:00:00Z",<br>        "status": "pending"<br>    },<br>    {<br>        "id": 0,<br>        "title": "Finish DSA",<br>        "description": "Finish on the BFS and DFS",<br>        "due_date": "2025-01-20T00:00:00Z",<br>        "status": "pending"<br>    }<br>] |
| code | 200 |
| cookie | [] |
| header | _complex array_ |
| name | Get Tasks |
| originalRequest | _object_ |
| status | OK |

####### item[1].response[0].header
| # | key| value|
| --- | --- | --- |
| 1 | Content-Type | application/json; charset=utf-8 |
| 2 | Date | Thu, 18 Dec 2025 11:00:58 GMT |
| 3 | Content-Length | 1125 |

####### item[1].response[0].originalRequest
| Property | Value |
| --- | --- |
| header | [] |
| method | GET |
| url | _object_ |

######## item[1].response[0].originalRequest.url
| Property | Value |
| --- | --- |
| path | [tasks] |
| raw | /tasks |

#### item[2]
| Property | Value |
| --- | --- |
| event | _complex array_ |
| name | Get Task |
| request | _object_ |
| response | _complex array_ |

##### item[2].event
| # | listen| script|
| --- | --- | --- |
| 1 | test | _complex_ |

###### item[2].event[0]
| Property | Value |
| --- | --- |
| listen | test |
| script | _object_ |

####### item[2].event[0].script
| Property | Value |
| --- | --- |
| exec | [pm.test('Get Task - Response has task details', function () {,     const jsonData = pm.response.json();,     pm.expect(jsonData).to.have.property('id');,     pm.expect(jsonData).to.have.property('title');, }), , pm.test('Get Task - ID matches requested ID', function () {,     const jsonData = pm.response.json();,     const requestedId = pm.collectionVariables.get('created_task_id');,     if (requestedId && jsonData.id) {,         pm.expect(String(jsonData.id)).to.eql(String(requestedId));,     }, }), , pm.test('Get Task - Status code is 200 OK', function () {,     pm.response.to.have.status(200);, })] |
| type | text/javascript |

##### item[2].request
| Property | Value |
| --- | --- |
| description | Retrieve a single task by its unique identifier.<br><br>**Endpoint**  <br>`GET {{base_url}}/tasks/:id`<br><br>In this request, the example URL `{{base_url}}/tasks/0` represents fetching the task with ID `0`. In practice, replace `0` with the ID of the task you want to retrieve.<br><br>**Path parameters**<br>- `id` (integer, required): Unique identifier of the task to retrieve.<br><br>**Authentication**<br>If your API is secured, ensure the appropriate auth is configured at the collection or request level (for example, a Bearer token using `{{token}}`). This request relies on that shared configuration and does not set auth explicitly.<br><br>**Successful response (200 OK)**<br>On success, the API returns a JSON object representing the task. At minimum, it includes:<br>- `id` (integer): ID of the task.  <br>- `title` (string): Short title of the task.<br><br>Depending on your implementation, the response may also include additional fields such as:<br>- `description` (string): Detailed description of the task.  <br>- `due_date` (string, date/time): When the task is due.  <br>- `status` (string): Current status of the task (for example, `pending`, `completed`).<br><br>This request has tests that validate:<br>- The response contains `id` and `title` properties.  <br>- The `id` in the response matches the `created_task_id` collection variable (if set).  <br>- The HTTP status code is `200`.  |
| header | [] |
| method | GET |
| url | _object_ |

###### item[2].request.url
| Property | Value |
| --- | --- |
| host | [{{base_url}}] |
| path | [tasks, 0] |
| raw | {{base_url}}/tasks/0 |

##### item[2].response
| # | _postman_previewlanguage| body| code| cookie| header| name| originalRequest| status|
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | null | {<br>    "id": 0,<br>    "title": "Finish report",<br>    "description": "Prepare Q1 report",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "completed"<br>} | 200 | [] | _complex_ | Get Task | _complex_ | OK |

###### item[2].response[0]
| Property | Value |
| --- | --- |
| _postman_previewlanguage | null |
| body | {<br>    "id": 0,<br>    "title": "Finish report",<br>    "description": "Prepare Q1 report",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "completed"<br>} |
| code | 200 |
| cookie | [] |
| header | _complex array_ |
| name | Get Task |
| originalRequest | _object_ |
| status | OK |

####### item[2].response[0].header
| # | key| value|
| --- | --- | --- |
| 1 | Content-Type | application/json; charset=utf-8 |
| 2 | Date | Thu, 18 Dec 2025 11:01:08 GMT |
| 3 | Content-Length | 152 |

####### item[2].response[0].originalRequest
| Property | Value |
| --- | --- |
| header | [] |
| method | GET |
| url | _object_ |

######## item[2].response[0].originalRequest.url
| Property | Value |
| --- | --- |
| path | [tasks, 0] |
| raw | /tasks/0 |

#### item[3]
| Property | Value |
| --- | --- |
| name | Delete Task |
| request | _object_ |
| response | _complex array_ |

##### item[3].request
| Property | Value |
| --- | --- |
| description | Deletes a task resource by its unique ID.<br><br>- **HTTP method**: `DELETE`<br>- **Endpoint**: `{{base_url}}/tasks/{id}`<br>- **Current example**: Uses a hard-coded ID `0` (`{{base_url}}/tasks/0`). Replace `0` with the actual task ID you want to delete.<br>- **Successful responses**:<br>  - `200 OK` – Backend returns a confirmation payload (e.g., the deleted task or a status message).<br>  - `204 No Content` – Backend indicates successful deletion without a response body.<br><br>Use this endpoint only when you intend to permanently remove the specified task. |
| header | [] |
| method | DELETE |
| url | _object_ |

###### item[3].request.url
| Property | Value |
| --- | --- |
| host | [{{base_url}}] |
| path | [tasks, 0] |
| raw | {{base_url}}/tasks/0 |

##### item[3].response
| # | _postman_previewlanguage| body| code| cookie| header| name| originalRequest| status|
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | null | {<br>    "message": "Task deleted successfully"<br>} | 200 | [] | _complex_ | Delete Task | _complex_ | OK |

###### item[3].response[0]
| Property | Value |
| --- | --- |
| _postman_previewlanguage | null |
| body | {<br>    "message": "Task deleted successfully"<br>} |
| code | 200 |
| cookie | [] |
| header | _complex array_ |
| name | Delete Task |
| originalRequest | _object_ |
| status | OK |

####### item[3].response[0].header
| # | key| value|
| --- | --- | --- |
| 1 | Content-Type | application/json; charset=utf-8 |
| 2 | Date | Thu, 18 Dec 2025 11:02:48 GMT |
| 3 | Content-Length | 46 |

####### item[3].response[0].originalRequest
| Property | Value |
| --- | --- |
| header | [] |
| method | DELETE |
| url | _object_ |

######## item[3].response[0].originalRequest.url
| Property | Value |
| --- | --- |
| path | [tasks, 0] |
| raw | /tasks/0 |

#### item[4]
| Property | Value |
| --- | --- |
| event | _complex array_ |
| name | Update Task |
| request | _object_ |
| response | _complex array_ |

##### item[4].event
| # | listen| script|
| --- | --- | --- |
| 1 | test | _complex_ |

###### item[4].event[0]
| Property | Value |
| --- | --- |
| listen | test |
| script | _object_ |

####### item[4].event[0].script
| Property | Value |
| --- | --- |
| exec | [pm.test('Update Task - Response confirms update', function () {,     const jsonData = pm.response.json();,     pm.expect(jsonData).to.have.property('id');, }), , pm.test('Update Task - Status code is 200 OK or 204 No Content', function () {,     pm.expect(pm.response.code).to.be.oneOf([,         200,,         204,     ]);, }), , pm.test('Update Task - Response has updated fields', function () {,     if (pm.response.code === 200) {,         const jsonData = pm.response.json();,         pm.expect(jsonData).to.have.property('title');,     }, })] |
| packages | _object_ |
| requests | _object_ |
| type | text/javascript |

######## item[4].event[0].script.packages

_Empty object_

######## item[4].event[0].script.requests

_Empty object_

##### item[4].request
| Property | Value |
| --- | --- |
| body | _object_ |
| description | Updates an existing task by its unique ID.<br><br>## Endpoint<br>- **Method:** `PUT`<br>- **URL:** `{{base_url}}/tasks/{id}`<br>- **Path parameter:**<br>  - `id` (integer, required) – ID of the task to update. In this request example, the value is `0`.<br><br>> Note: The request URL uses the `{{base_url}}` variable, which should be defined in your active environment.<br><br>## Request Body<br>JSON object with the following fields:<br><br>```json<br>{<br>  "id": 0,<br>  "title": "Finish DSA",<br>  "description": "Finish on the BFS and DFS",<br>  "due_date": "2025-01-20T00:00:00Z",<br>  "status": "completed"<br>}<br>```<br><br>### Body schema<br>- `id` (integer, required)<br>  - Unique identifier of the task. Should match the `id` in the path.<br>- `title` (string, required)<br>  - Short, human-readable title of the task.<br>- `description` (string, optional)<br>  - Detailed description of the task.<br>- `due_date` (string, required)<br>  - Due date/time of the task in ISO 8601 format (e.g., `2025-01-20T00:00:00Z`).<br>- `status` (string, required)<br>  - Current status of the task, e.g., `pending`, `in_progress`, `completed`.<br><br>## Responses<br>- **200 OK** – Task was successfully updated and the updated task object is returned in the response body.<br>- **204 No Content** – Task was successfully updated and no response body is returned.<br>- **400 Bad Request** – The request body is invalid (e.g., missing required fields, invalid data types, or mismatched `id`).<br>- **404 Not Found** – No task exists with the specified `id`.<br>- **500 Internal Server Error** – An unexpected error occurred on the server while processing the update. |
| header | [] |
| method | PUT |
| url | _object_ |

###### item[4].request.body
| Property | Value |
| --- | --- |
| mode | raw |
| options | _object_ |
| raw | {<br>    "id": 0,<br>    "title": "Finish DSA",<br>    "description": "Finish on the BFS and DFS",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "completed"<br>} |

####### item[4].request.body.options
| Property | Value |
| --- | --- |
| raw | _object_ |

######## item[4].request.body.options.raw
| Property | Value |
| --- | --- |
| language | json |

###### item[4].request.url
| Property | Value |
| --- | --- |
| host | [{{base_url}}] |
| path | [tasks, 0] |
| raw | {{base_url}}/tasks/0 |

##### item[4].response
| # | _postman_previewlanguage| body| code| cookie| header| name| originalRequest| status|
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | null | {<br>    "id": 0,<br>    "title": "Finish DSA",<br>    "description": "Finish on the BFS and DFS",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "completed"<br>} | 200 | [] | _complex_ | Update Task | _complex_ | OK |

###### item[4].response[0]
| Property | Value |
| --- | --- |
| _postman_previewlanguage | null |
| body | {<br>    "id": 0,<br>    "title": "Finish DSA",<br>    "description": "Finish on the BFS and DFS",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "completed"<br>} |
| code | 200 |
| cookie | [] |
| header | _complex array_ |
| name | Update Task |
| originalRequest | _object_ |
| status | OK |

####### item[4].response[0].header
| # | key| value|
| --- | --- | --- |
| 1 | Content-Type | application/json; charset=utf-8 |
| 2 | Date | Thu, 18 Dec 2025 11:02:40 GMT |
| 3 | Content-Length | 157 |

####### item[4].response[0].originalRequest
| Property | Value |
| --- | --- |
| body | _object_ |
| header | [] |
| method | PUT |
| url | _object_ |

######## item[4].response[0].originalRequest.body
| Property | Value |
| --- | --- |
| mode | raw |
| options | _object_ |
| raw | {<br>    "id": 0,<br>    "title": "Finish DSA",<br>    "description": "Finish on the BFS and DFS",<br>    "due_date": "2025-01-20T00:00:00Z",<br>    "status": "completed"<br>} |

######### item[4].response[0].originalRequest.body.options
| Property | Value |
| --- | --- |
| raw | _object_ |

########## item[4].response[0].originalRequest.body.options.raw
| Property | Value |
| --- | --- |
| language | json |

######## item[4].response[0].originalRequest.url
| Property | Value |
| --- | --- |
| path | [tasks, 0] |
| raw | /tasks/0 |