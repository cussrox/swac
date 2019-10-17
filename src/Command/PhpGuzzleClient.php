<?php

namespace Swac\Command;

use Swac\Log;
use Swac\Php\Client\Client;
use Swac\Rest\Rest;
use Swac\Swagger\Reader;
use Symfony\Component\Yaml\Yaml;
use Yaoi\Command;
use Yaoi\Command\Definition;

class PhpGuzzleClient extends Command
{
    public $namespace;
    public $schemaPath;
    public $projectPath = './';
    public $operations;

    /**
     * @param Definition $definition
     * @param \stdClass|static $options
     */
    static function setUpDefinition(Definition $definition, $options)
    {
        $options->schemaPath = Command\Option::create()->setType()->setIsRequired()->setIsUnnamed()
            ->setDescription('Path to swagger.json');
        $options->projectPath = Command\Option::create()->setType()
            ->setDescription('Path to project root, default ./');
        $options->namespace = Command\Option::create()->setType()->setIsRequired()
            ->setDescription('Project namespace');
        $options->operations = Command\Option::create()->setType()
            ->setDescription('Operations filter in form of comma-separated list of method/path, default empty');
    }

    public function performAction()
    {
        $rest = new Rest();

        $phpClient = new Client($this->namespace, $this->projectPath);
        $rest->addRenderer($phpClient);

        $reader = new Reader($rest);
        $reader->setLog(Log::getInstance());
        $schemaJson = json_decode(file_get_contents($this->schemaPath));

        $reader->addSchemaJson($schemaJson);
        $reader->process();

        $phpClient->store($this->projectPath);
    }
}