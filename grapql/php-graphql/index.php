<?php

require_once 'vendor/autoload.php';
require_once 'db.php';

use GraphQL\Type\Definition\ObjectType;
use GraphQL\Type\Definition\Type;
use GraphQL\Type\Schema;
use GraphQL\GraphQL;


// define Type
$queryType = new ObjectType([
    'name' => 'Query',
    'fields' => [
        'hello' => [
            'type' => Type::string(),
            'resolve' => function ($root, $args) {
                return 'Hello World!';
            }
        ],        
        'ping' => [
            'type' => Type::int(),
            'resolve' => function ($root, $args) {
                return 69;
            }
        ],
        'book' => [
            'type' => new ObjectType([
                'name' => 'Book',
                'fields' => [
                    'id' => Type::int(),
                    'title' => Type::string(),
                    'author' => Type::string(),
                ]
            ]),
            'args' => [
                'id' => Type::int(),
            ],
            'resolve' => function ($root, $args) use ($pdo) {
                $stmt = $pdo->prepare('SELECT * FROM books WHERE id = :id');
                $stmt->execute(['id' => $args['id']]);
                return $stmt->fetch();
            }
        ]
    ]
]);


// Setup schema
$schema = new Schema([
    'query' => $queryType
]);

// Coding Handler
$requestPayload = isset($_GET['query']) ? $_GET['query'] : null;
if (!$requestPayload && isset($_POST['query'])) {
    $requestPayload = $_POST['query'];
    var_dump( $requestPayload );
}

try {
    $result = GraphQL::executeQuery($schema, $requestPayload);
    $output = $result->toArray();
} catch (\Exception $e) {
    $output = [
        'errors' => [
            [
                'message' => $e->getMessage()
            ]
        ]
    ];
}

// respond
header('Content-Type: application/json');
echo json_encode($output);