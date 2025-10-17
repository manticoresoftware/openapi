<?php
/**
 * MySQL Variables Extractor
 * 
 * This script connects to a MySQL database and extracts all system variables
 * into a PHP array. It supports both global and session variables.
 */

class MySQLVariablesExtractor {
    private $connection;
    private $host;
    private $username;
    private $password;
    private $database;
    private $port;
    
    public function __construct($host = 'localhost', $username = 'root', $password = '', $database = '', $port = 3306) {
        $this->host = $host;
        $this->username = $username;
        $this->password = $password;
        $this->database = $database;
        $this->port = $port;
    }
    
    /**
     * Connect to MySQL database
     */
    public function connect() {
        try {
            $this->connection = mysqli_connect($this->host, $this->username, $this->password, $this->database, $this->port);
            
            if (!$this->connection) {
                throw new Exception("Connection failed: " . mysqli_connect_error());
            }
            
            // Set character set
            mysqli_set_charset($this->connection, "utf8mb4");
            
            return true;
        } catch (Exception $e) {
            throw new Exception("Connection failed: " . $e->getMessage());
        }
    }
    
    /**
     * Extract all variables with default values using a single query
     */
    public function getAllVariables() {
        $variables = [];
        
        try {
            // Get all variables in a single query
            $result = mysqli_query($this->connection, "SHOW VARIABLES");
            if (!$result) {
                throw new Exception("Failed to get variables: " . mysqli_error($this->connection));
            }
            
            while ($row = mysqli_fetch_assoc($result)) {
                $variables[$row['Variable_name']] = [
                    'current' => $row['Value'],
                    'default' => null
                ];
            }
            
            mysqli_free_result($result);
            
            // Get default values from information_schema
            $result = mysqli_query($this->connection, "
                SELECT VARIABLE_NAME, VARIABLE_DEFAULT 
                FROM information_schema.SESSION_VARIABLES 
                WHERE VARIABLE_DEFAULT IS NOT NULL
            ");
            
            if ($result) {
                while ($row = mysqli_fetch_assoc($result)) {
                    $varName = $row['VARIABLE_NAME'];
                    if (isset($variables[$varName])) {
                        $variables[$varName]['default'] = $row['VARIABLE_DEFAULT'];
                    }
                }
                mysqli_free_result($result);
            }
            
        } catch (Exception $e) {
            throw new Exception("Failed to get variables: " . $e->getMessage());
        }
        
        return $variables;
    }
    
    /**
     * Extract global variables only
     */
    public function getGlobalVariables() {
        $variables = [];
        
        try {
            $result = mysqli_query($this->connection, "SHOW GLOBAL VARIABLES");
            if (!$result) {
                throw new Exception("Failed to get global variables: " . mysqli_error($this->connection));
            }
            
            while ($row = mysqli_fetch_assoc($result)) {
                $variables[$row['Variable_name']] = [
                    'current' => $row['Value'],
                    'default' => null
                ];
            }
            
            mysqli_free_result($result);
            
            // Get default values from information_schema
            $result = mysqli_query($this->connection, "
                SELECT VARIABLE_NAME, VARIABLE_DEFAULT 
                FROM information_schema.GLOBAL_VARIABLES 
                WHERE VARIABLE_DEFAULT IS NOT NULL
            ");
            
            if ($result) {
                while ($row = mysqli_fetch_assoc($result)) {
                    $varName = $row['VARIABLE_NAME'];
                    if (isset($variables[$varName])) {
                        $variables[$varName]['default'] = $row['VARIABLE_DEFAULT'];
                    }
                }
                mysqli_free_result($result);
            }
            
        } catch (Exception $e) {
            throw new Exception("Failed to get global variables: " . $e->getMessage());
        }
        
        return $variables;
    }
    
    /**
     * Extract session variables only
     */
    public function getSessionVariables() {
        $variables = [];
        
        try {
            $result = mysqli_query($this->connection, "SHOW SESSION VARIABLES");
            if (!$result) {
                throw new Exception("Failed to get session variables: " . mysqli_error($this->connection));
            }
            
            while ($row = mysqli_fetch_assoc($result)) {
                $variables[$row['Variable_name']] = [
                    'current' => $row['Value'],
                    'default' => null
                ];
            }
            
            mysqli_free_result($result);
            
            // Get default values from information_schema
            $result = mysqli_query($this->connection, "
                SELECT VARIABLE_NAME, VARIABLE_DEFAULT 
                FROM information_schema.SESSION_VARIABLES 
                WHERE VARIABLE_DEFAULT IS NOT NULL
            ");
            
            if ($result) {
                while ($row = mysqli_fetch_assoc($result)) {
                    $varName = $row['VARIABLE_NAME'];
                    if (isset($variables[$varName])) {
                        $variables[$varName]['default'] = $row['VARIABLE_DEFAULT'];
                    }
                }
                mysqli_free_result($result);
            }
            
        } catch (Exception $e) {
            throw new Exception("Failed to get session variables: " . $e->getMessage());
        }
        
        return $variables;
    }
    
    /**
     * Get MySQL version and server information
     */
    public function getServerInfo() {
        $info = [];
        
        try {
            // Get MySQL version
            $result = mysqli_query($this->connection, "SELECT VERSION() as version");
            if (!$result) {
                throw new Exception("Failed to get version: " . mysqli_error($this->connection));
            }
            $row = mysqli_fetch_assoc($result);
            $info['version'] = $row['version'];
            mysqli_free_result($result);
            
            // Get server status
            $result = mysqli_query($this->connection, "SHOW STATUS LIKE 'Uptime'");
            if (!$result) {
                throw new Exception("Failed to get uptime: " . mysqli_error($this->connection));
            }
            $row = mysqli_fetch_assoc($result);
            $info['uptime'] = $row['Value'];
            mysqli_free_result($result);
            
            // Get current database
            $result = mysqli_query($this->connection, "SELECT DATABASE() as current_database");
            if (!$result) {
                throw new Exception("Failed to get current database: " . mysqli_error($this->connection));
            }
            $row = mysqli_fetch_assoc($result);
            $info['current_database'] = $row['current_database'];
            mysqli_free_result($result);
            
        } catch (Exception $e) {
            throw new Exception("Failed to get server info: " . $e->getMessage());
        }
        
        return $info;
    }
    
    /**
     * Close the database connection
     */
    public function close() {
        if ($this->connection) {
            mysqli_close($this->connection);
            $this->connection = null;
        }
    }
    
    /**
     * Export variables to a PHP file
     */
    public function exportToFile($filename = 'mysql_variables.php') {
        $variables = $this->getAllVariables();
        $serverInfo = $this->getServerInfo();
        
        $content = "<?php\n";
        $content .= "/**\n";
        $content .= " * MySQL Variables Export\n";
        $content .= " * Generated on: " . date('Y-m-d H:i:s') . "\n";
        $content .= " * MySQL Version: " . $serverInfo['version'] . "\n";
        $content .= " */\n\n";
        $content .= "\$mysql_variables = " . var_export($variables, true) . ";\n\n";
        $content .= "\$mysql_server_info = " . var_export($serverInfo, true) . ";\n\n";
        $content .= "// Access variables like: \$mysql_variables['max_connections']['current']\n";
        $content .= "// Access server info like: \$mysql_server_info['version']\n";
        
        if (file_put_contents($filename, $content)) {
            return "Variables exported to {$filename}";
        } else {
            throw new Exception("Failed to write to file: {$filename}");
        }
    }
}

// Usage example
try {
    // Configuration - modify these values as needed
    $config = [
        'host' => 'localhost',
        'username' => 'root',
        'password' => '',
        'database' => '',
        'port' => 3306
    ];
    
    // Create extractor instance
    $extractor = new MySQLVariablesExtractor(
        $config['host'],
        $config['username'],
        $config['password'],
        $config['database'],
        $config['port']
    );
    
    // Connect to database
    $extractor->connect();
    echo "Connected to MySQL successfully!\n";
    
    // Get server information
    $serverInfo = $extractor->getServerInfo();
    echo "MySQL Version: " . $serverInfo['version'] . "\n";
    echo "Server Uptime: " . $serverInfo['uptime'] . " seconds\n";
    echo "Current Database: " . ($serverInfo['current_database'] ?: 'None') . "\n\n";
    
            // Extract all variables
        $allVariables = $extractor->getAllVariables();
        
        // Display summary
        echo "Found " . count($allVariables) . " variables\n\n";
    
    // Export to file
    $result = $extractor->exportToFile('mysql_variables.php');
    echo $result . "\n";
    
    // Display some important variables
    echo "\nImportant Variables:\n";
    $importantVars = [
        'max_connections', 'innodb_buffer_pool_size', 'max_allowed_packet',
        'wait_timeout', 'interactive_timeout', 'character_set_server',
        'collation_server', 'sql_mode', 'autocommit'
    ];
    
    foreach ($importantVars as $var) {
        if (isset($allVariables[$var])) {
            $current = $allVariables[$var]['current'];
            $default = $allVariables[$var]['default'];
            $defaultInfo = $default !== null ? " (default: $default)" : " (no default)";
            echo "  {$var}: {$current}{$defaultInfo}\n";
        }
    }
    
    // Close connection
    $extractor->close();
    
} catch (Exception $e) {
    echo "Error: " . $e->getMessage() . "\n";
    exit(1);
}
?> 