import requests
import json

from bravado.client import SwaggerClient
from bravado.exception import HTTPError
from bravado.requests_client import RequestsClient
from bravado.swagger_model import load_url

http_client = RequestsClient()
spec_dict = load_url('http://localhost:8092/v1/swagger.json', http_client=http_client)
spec_dict['host'] = '{}:{}'.format("localhost", 8092)
client = SwaggerClient.from_spec(spec_dict, config={'use_models': False})


def do_request(func, **kwargs):
    request = func(
        **kwargs
    )
    try:
        response = request.response()

    except HTTPError as e:
        print(f"got error while do request with status_code={ e.response.status_code } and body={ e.response.text }")
        return None, {"status_code": e.response.status_code, "body": json.loads(e.response.text)}
    else:
        return response, None
