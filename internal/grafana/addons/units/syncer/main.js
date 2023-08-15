const path = require('path');
const fs = require('fs');
const Parser = require('tree-sitter');
const JavaScript = require('tree-sitter-javascript');

const parser = new Parser();
parser.setLanguage(JavaScript);

// list all files in `definitions` folder
// for each file, parse the file and extract the unit name and to_anchor value
// write the result to a file
function readFiles() {
    const result = [];
    fs.readdirSync(path.join(__dirname, 'definitions')).forEach((file) => {
        result.push(readFile(file));
    });
    console.log(JSON.stringify(result, null, 2));
}

function readFile(file) {
    const sourceCode = fs.readFileSync(path.join(__dirname, 'definitions', file), 'utf8');
    const tree = parser.parse(sourceCode);
    const query = readQuery();

    const result = {
        unit: path.basename(file, '.js'),
        formats: [],
    };

    query.captures(tree.rootNode).forEach((capture) => {
        var id;
        switch (capture.name) {
            case 'id':
                switch (capture.node.type) {
                    case 'property_identifier':
                        id = capture.node.text;
                        break;
                    case 'string':
                        id = capture.node.text.slice(1, -1);
                        break;
                    default:
                        console.error('Unknown type: ' + capture.node.type);
                }
                result.formats.push({ id, ref: "" });
                break;
            default:
                break;
        }
    });

    return result;
}

function readQuery() {
    const query = fs.readFileSync(path.join(__dirname, 'query.scm'), 'utf8');
    return new Parser.Query(JavaScript, query)
}

readFiles();
