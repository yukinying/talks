    connection.query(
        'SELECT * FROM users WHERE id = ' + userId,  // HL
        function(err, results) {  /* Call me Bobby Tables */});
