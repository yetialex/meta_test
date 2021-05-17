from common import (
    client,
    do_request
)


def test_directories():
    directory = {
        "name": "zyfra",
        "description": "zyfra real time meta",
    }

    resp, error = do_request(client.core.getDirectory, directory_id=100)
    assert error['status_code'] == 404

    resp, error = do_request(client.core.createDirectory, body=directory)
    assert error is None
    directory_id = resp.result['id']

    resp, error = do_request(client.core.getDirectory, directory_id=directory_id)
    assert error is None

    resp, error = do_request(client.core.updateDirectory, directory_id=directory_id, body={"name": "zyfra_batch"})
    assert error is None

    resp, error = do_request(client.core.updateDirectory, directory_id=directory_id, body={"description": "zyfra batch"})
    assert error is None

    resp, error = do_request(client.core.updateDirectory, directory_id=directory_id, body={})
    assert error['status_code'] == 400

    resp, error = do_request(client.core.deleteDirectory, directory_id=directory_id)
    assert error is None

