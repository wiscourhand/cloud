namespace go demo



struct IdlInfo {
    1: string content(go.tag = 'json:"content"'),
    2: map<string, string> includes(go.tag = 'json:"includes"'),
}

struct IdlReq {
    1: string idlName(api.body='name'),
    2: IdlInfo idlInfo(api.body='info'),
}

struct AddIdlResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

service IdlService {
    AddIdlResp Register(1: IdlReq idlReq)(api.post = '/add-ldl-info')
    IdlInfo Query(1: string name)(api.get = '/query')
}
