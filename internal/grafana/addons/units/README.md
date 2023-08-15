# Unit Converter from Grafana to Guance

## Quick view

In Grafana and Guance Dashboard, there are many units to represent to data semantic.

This package is a converter to convert the unit string from Grafana to Guance.

The implementation steps are:

1. Get the full list of units from _Grafana_.
2. Extract all the units from source code of _Guance Cloud Console_ by tree-sitter query.
3. Compare the units from Grafana and Guance, and find the mapping relationship.

## Extract units from Guance

The Guance Cloud Console is written in React, and the source code is compiled to JavaScript.

The tree-sitter is a parser generator tool and incremental parsing library. It can extract any information from source code and its CST (Concrete Syntax Tree).

The query as follows:

```lisp
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
```

The steps to extract the units from Guance:

1. Clone the frontend source code of Guance Cloud Console.
2. cd `lib/convert-units/lib/definitions` folder.
3. Run the script `node main.js`.
