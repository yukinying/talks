    connection.query(
        'SELECT * FROM users WHERE id = ?', [userId],  // HL
        function(err, results) { /* use sql prepared statement */ });
