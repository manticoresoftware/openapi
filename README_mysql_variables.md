# MySQL Variables Extractor

This set of PHP scripts allows you to extract all MySQL system variables into a PHP array for analysis, monitoring, or configuration management.

## Files

- `extract_mysql_variables.php` - Main script that connects to MySQL and extracts variables
- `example_usage.php` - Example script showing how to use the extracted variables
- `mysql_variables.php` - Generated file containing all MySQL variables (created after running the extractor)

## Features

- Extracts both global and session variables
- Includes server information (version, uptime, current database)
- Exports variables to a PHP file for easy inclusion in other scripts
- Provides command-line and web interface
- Includes error handling and connection management
- Shows important variables with descriptions

## Requirements

- PHP 7.0 or higher
- PDO MySQL extension
- Access to a MySQL server

## Installation

1. Make sure you have PHP and the PDO MySQL extension installed
2. Place the scripts in your web directory or run from command line
3. Configure the database connection settings in `extract_mysql_variables.php`

## Usage

### Command Line

```bash
php extract_mysql_variables.php
```

### Web Browser

1. Open `extract_mysql_variables.php` in your web browser
2. Click the "Run Extractor" button
3. The script will generate `mysql_variables.php` with all variables

### Configuration

Edit the configuration array in `extract_mysql_variables.php`:

```php
$config = [
    'host' => 'localhost',
    'username' => 'root',
    'password' => '',
    'database' => '',
    'port' => 3306
];
```

## Output

The script generates a `mysql_variables.php` file containing:

```php
<?php
/**
 * MySQL Variables Export
 * Generated on: 2024-01-15 10:30:00
 * MySQL Version: 8.0.35
 */

$mysql_variables = [
    'max_connections' => [
        'current' => '151',
        'default' => '151'
    ],
    'innodb_buffer_pool_size' => [
        'current' => '134217728',
        'default' => '134217728'
    ],
    'character_set_client' => [
        'current' => 'utf8mb4',
        'default' => 'utf8mb4'
    ],
    'time_zone' => [
        'current' => 'SYSTEM',
        'default' => 'SYSTEM'
    ],
    // ... all variables with current and default values
];

$mysql_server_info = [
    'version' => '8.0.35',
    'uptime' => '12345',
    'current_database' => 'test'
];
```

## Using the Extracted Variables

### Include the generated file

```php
require_once 'mysql_variables.php';
```

### Access variables

```php
// Access current values
$maxConnections = $mysql_variables['max_connections']['current'];
$bufferPoolSize = $mysql_variables['innodb_buffer_pool_size']['current'];
$characterSet = $mysql_variables['character_set_client']['current'];
$timezone = $mysql_variables['time_zone']['current'];

// Access default values
$maxConnectionsDefault = $mysql_variables['max_connections']['default'];

// Server information
$version = $mysql_server_info['version'];
$uptime = $mysql_server_info['uptime'];
```

### Example Analysis

Run the example script to see how to analyze the variables:

```bash
php example_usage.php
```

This will show:
- Server information
- Important global and session variables
- Performance analysis and recommendations
- Variable statistics and categorization

## Important Variables

The script highlights these important MySQL variables:

### Global Variables
- `max_connections` - Maximum number of connections
- `innodb_buffer_pool_size` - InnoDB buffer pool size
- `max_allowed_packet` - Maximum allowed packet size
- `wait_timeout` - Wait timeout in seconds
- `interactive_timeout` - Interactive timeout in seconds
- `character_set_server` - Server character set
- `collation_server` - Server collation
- `sql_mode` - SQL mode
- `autocommit` - Autocommit mode

### Session Variables
- `character_set_client` - Client character set
- `character_set_connection` - Connection character set
- `character_set_database` - Database character set
- `character_set_results` - Results character set
- `time_zone` - Session timezone
- `sql_mode` - Session SQL mode
- `autocommit` - Session autocommit mode

## Security Considerations

- Store database credentials securely
- Use a dedicated MySQL user with limited privileges
- Consider using environment variables for sensitive data
- Don't expose the generated variables file publicly

## Troubleshooting

### Connection Issues
- Verify MySQL server is running
- Check host, port, username, and password
- Ensure the user has privileges to view variables

### Permission Issues
- Make sure the script has write permissions to create the output file
- Check that the web server can write to the directory

### Missing Variables
- Some variables may not be available depending on MySQL version
- Check MySQL documentation for version-specific variables

## Advanced Usage

### Custom Variable Filtering

```php
// Get only InnoDB-related variables
$innodbVars = [];
foreach ($mysql_variables as $var => $data) {
    if (strpos($var, 'innodb_') === 0) {
        $innodbVars[$var] = $data;
    }
}
```

### Variable Comparison

```php
// Compare global vs session variables (current values only)
$globalCurrent = array_map(function($data) { return $data['current']; }, $mysql_variables['global']);
$sessionCurrent = array_map(function($data) { return $data['current']; }, $mysql_variables['session']);
$differences = array_diff_assoc($globalCurrent, $sessionCurrent);
```

### Performance Monitoring

```php
// Monitor specific performance variables
$performanceVars = [
    'innodb_buffer_pool_size',
    'query_cache_size',
    'tmp_table_size',
    'max_heap_table_size'
];

foreach ($performanceVars as $var) {
    if (isset($mysql_variables['global'][$var])) {
        $current = $mysql_variables['global'][$var]['current'];
        $default = $mysql_variables['global'][$var]['default'];
        echo "$var: $current (default: " . ($default ?? 'none') . ")\n";
    }
}
```

## License

This script is provided as-is for educational and development purposes. Use at your own risk in production environments. 