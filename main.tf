provider "solace" {
    host = "localhost:8080"
    base_path = "/SEMP/v2/config"
    user = "admin"
    password = "alex"
    msg_vpn = "go-tf-1"
}

resource "solace_msgvpn" "my-vpn" {
    name = "go-tf-1"
    enabled = true
    authentication_basic_enabled = false
    max_spool_usage = 1100
	max_connection_count = 123
	max_egress_flow_count = 456
	max_endpoint_count = 789
	max_ingress_flow_count = 111
	max_spool_usage = 222
	max_subscription_count = 333
	max_transacted_session_count = 444
	max_transaction_count = 555
}

resource "solace_aclprofile" "my-acl" {
    name = "ach-test-acl-1"
    client_connection_default_action = "disallow"
    publish_topic_default_action = "allow"
    subscribe_topic_default_action = "disallow"
}

resource "solace_aclprofile_clientconnexception" "my-machine-exc" {
    acl = "ach-test-acl-1"
    address = "127.0.0.1/10"
}

resource "solace_aclprofile_publishexception" "my-mon-allow" {
    acl = "ach-test-acl-1"
    topic_syntax = "smf"
    topic = "box-foobar/>"
}

resource "solace_aclprofile_subscribeexception" "my-sub-allow" {
    acl = "ach-test-acl-1"
    topic_syntax = "smf"
    topic = "box3-foobar/>"
}

resource "solace_aclprofile_subscribeexception" "my-imported-sub" {
    acl = "ach-test-acl-1"
    topic = "box-import-topi/>"
    topic_syntax = "smf"
    
}

resource "solace_aclprofile_subscribeexception" "my-sub-mqtt-allow" {
    acl = "ach-test-acl-1"
    topic_syntax = "mqtt"
    topic = "box3-mqtt-foobar/>"
}