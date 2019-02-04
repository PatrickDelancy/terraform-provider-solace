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

# ACL profiles
resource "solace_aclprofile" "my-acl" {
    name = "ach-test-acl-1"
    msg_vpn = "${solace_msgvpn.my-vpn.name}"
    client_connection_default_action = "disallow"
    publish_topic_default_action = "allow"
    subscribe_topic_default_action = "disallow"
}

resource "solace_aclprofile_clientconnexception" "my-machine-exc" {
    acl = "${solace_aclprofile.my-acl.name}"
    address = "127.0.0.1/10"
}

resource "solace_aclprofile_publishexception" "my-mon-allow" {
    acl = "${solace_aclprofile.my-acl.name}"
    topic_syntax = "smf"
    topic = "box-foobar/>"
}

resource "solace_aclprofile_subscribeexception" "my-sub-allow" {
    acl = "${solace_aclprofile.my-acl.name}"
    topic_syntax = "smf"
    topic = "box3-foobar/>"
}

resource "solace_aclprofile_subscribeexception" "my-imported-sub" {
    acl = "${solace_aclprofile.my-acl.name}"
    topic = "box-import-topi/>"
    topic_syntax = "smf"
    
}

resource "solace_aclprofile_subscribeexception" "my-sub-mqtt-allow" {
    acl = "${solace_aclprofile.my-acl.name}"
    topic_syntax = "mqtt"
    topic = "box3-mqtt-foobar/>"
}

# Client profiles
resource "solace_clientprofile" "my-client-profile" {
    name = "ach-client-prof"
    msg_vpn = "${solace_msgvpn.my-vpn.name}"
    allow_cut_through_forwarding = true
    allow_guaranteed_endpoint_create = true
    allow_guaranteed_msg_receive = false
    allow_guaranteed_msg_send = false
    allow_transacted_sessions = false
}

# Client username
resource "solace_clientusername" "my-client-name" {
    name = "ach-client-name"
    acl = "${solace_aclprofile.my-acl.name}"
    msg_vpn = "${solace_msgvpn.my-vpn.name}"
    profile = "${solace_clientprofile.my-client-profile.name}"
    enabled = false
    password = "ach-test"
}
