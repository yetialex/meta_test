import random
from common import (
    client,
    do_request
)


def test_scenario1():

    gate_name = f"p3apr3-pda-vc.p3.ia.{random.randint(1, 100000)}.nlmk"
    mnt = f"dc2-dp7iba-{random.randint(1, 1000)}-s"

    # clean up
    resp, error = do_request(client.iba.deleteIBAGateByName, gate_name=gate_name)

    # main pipeline
    resp, error = do_request(client.iba.getGateSourceByGateNameAndMnt, gate_name=gate_name, mnt=mnt)
    assert error['status_code'] == 404

    resp, error = do_request(client.iba.getIBAGateByName, gate_name=gate_name)
    assert error['status_code'] == 404

    gate = {
        "name": gate_name,
    }
    resp, error = do_request(client.iba.createOrUpdateIBAGate, body=gate)
    assert error is None
    gate_id = resp.result['id']

    gate = {
        "name": gate_name,
        "comment": "comment",
    }
    resp, error = do_request(client.iba.createOrUpdateIBAGate, body=gate)
    assert error is None

    gate = {
        "id": gate_id+10,
        "name": gate_name,
        "comment": "comment2",
    }
    resp, error = do_request(client.iba.createOrUpdateIBAGate, body=gate)
    assert error['status_code'] == 400

    resp, error = do_request(client.iba.createOrUpdateGateSourceByGateNameAndMnt, gate_name=gate_name, mnt=mnt, body={})
    assert error is None

    # clean up
    # resp, error = do_request(client.iba.deleteGateSourceByGateNameAndMnt, gate_name=gate_name, mnt=mnt)
    # assert error is None
    #
    # resp, error = do_request(client.iba.deleteIBAGateByName, gate_name=gate_name)
    # assert error is None

