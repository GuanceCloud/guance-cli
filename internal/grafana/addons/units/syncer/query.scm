; extract the unit formats
(program
    (expression_statement 
        (assignment_expression left: 
            (identifier) right: 
            (object [
                (pair key: (string) @id)
                (pair key: (property_identifier) @id)
            ])
        )
    )
)
