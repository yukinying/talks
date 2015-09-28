func consume(){
    // sample form of message: {"response":{"headers":{"X-Frame-Options":["SAMEORIGIN"]...
    //                         {"action":"element.triggered","events":["click"],"keyCha...
    type message struct { // HL
        *responseMessage 
        *domMessage      
    }
    type responseMessage struct { // HL 
        Response response       // with Json field "response"
    }
    type domMessage struct { // HL 
        Action   string         // with Json field "action"
    }
    dec := json.NewDecoder(stdout)    
    for {
        var m message
        err := dec.Decode(&m)  // HL
        if m.responseMessage != nil {
            //...
        } else if m.domMessage != nil {
            //...
        }
    }
}