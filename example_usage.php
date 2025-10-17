<?php
/**
 * Example Usage of Extracted MySQL Variables
 * 
 * This script demonstrates how to use the MySQL variables that were
 * extracted by the extract_mysql_variables.php script.
 */

// Include the generated variables file
if (file_exists('mysql_variables.php')) {
    require_once 'mysql_variables.php';
    
    echo "=== MySQL Variables Analysis ===\n\n";
    
    // Display server information
    echo "Server Information:\n";
    echo "MySQL Version: " . $mysql_server_info['version'] . "\n";
    echo "Server Uptime: " . $mysql_server_info['uptime'] . " seconds\n";
    echo "Current Database: " . ($mysql_server_info['current_database'] ?: 'None') . "\n\n";
    
    // Display important variables
    echo "Important Variables:\n";
    $importantVars = [
        'max_connections' => 'Maximum number of connections',
        'innodb_buffer_pool_size' => 'InnoDB buffer pool size',
        'max_allowed_packet' => 'Maximum allowed packet size',
        'wait_timeout' => 'Wait timeout in seconds',
        'interactive_timeout' => 'Interactive timeout in seconds',
        'character_set_server' => 'Server character set',
        'collation_server' => 'Server collation',
        'sql_mode' => 'SQL mode',
        'autocommit' => 'Autocommit mode',
        'innodb_log_file_size' => 'InnoDB log file size',
        'innodb_log_buffer_size' => 'InnoDB log buffer size',
        'query_cache_size' => 'Query cache size',
        'tmp_table_size' => 'Temporary table size',
        'max_heap_table_size' => 'Maximum heap table size',
        'character_set_client' => 'Client character set',
        'character_set_connection' => 'Connection character set',
        'character_set_database' => 'Database character set',
        'character_set_results' => 'Results character set',
        'time_zone' => 'Session timezone'
    ];
    
    foreach ($importantVars as $var => $description) {
        if (isset($mysql_variables[$var])) {
            $current = $mysql_variables[$var]['current'];
            $default = $mysql_variables[$var]['default'];
            $defaultInfo = $default !== null ? " (default: $default)" : " (no default)";
            echo sprintf("  %-25s: %s%s (%s)\n", $var, $current, $defaultInfo, $description);
        }
    }
    
    echo "\n";
    
    // Performance analysis
    echo "Performance Analysis:\n";
    
    // Check buffer pool size
    if (isset($mysql_variables['innodb_buffer_pool_size'])) {
        $bufferPoolSize = $mysql_variables['innodb_buffer_pool_size']['current'];
        if (is_numeric($bufferPoolSize)) {
            $bufferPoolMB = $bufferPoolSize / 1024 / 1024;
            echo "  InnoDB Buffer Pool Size: " . number_format($bufferPoolMB, 2) . " MB\n";
            
            if ($bufferPoolMB < 256) {
                echo "    ⚠️  Consider increasing buffer pool size for better performance\n";
            } elseif ($bufferPoolMB > 2048) {
                echo "    ✅ Buffer pool size looks good\n";
            }
        }
    }
    
    // Check max connections
    if (isset($mysql_variables['max_connections'])) {
        $maxConnections = $mysql_variables['max_connections']['current'];
        echo "  Max Connections: $maxConnections\n";
        
        if ($maxConnections < 100) {
            echo "    ⚠️  Consider increasing max_connections for high-traffic applications\n";
        } elseif ($maxConnections > 1000) {
            echo "    ⚠️  High max_connections may consume more memory\n";
        } else {
            echo "    ✅ Max connections setting looks reasonable\n";
        }
    }
    
    // Check wait timeout
    if (isset($mysql_variables['wait_timeout'])) {
        $waitTimeout = $mysql_variables['wait_timeout']['current'];
        echo "  Wait Timeout: $waitTimeout seconds\n";
        
        if ($waitTimeout < 60) {
            echo "    ⚠️  Very short wait timeout may cause connection issues\n";
        } elseif ($waitTimeout > 28800) {
            echo "    ⚠️  Very long wait timeout may waste resources\n";
        } else {
            echo "    ✅ Wait timeout setting looks reasonable\n";
        }
    }
    
    echo "\n";
    
    // Search for specific variables
    echo "Searching for variables containing 'cache':\n";
    $cacheVars = [];
    foreach ($mysql_variables as $var => $data) {
        if (stripos($var, 'cache') !== false) {
            $cacheVars[$var] = $data;
        }
    }
    
    if (!empty($cacheVars)) {
        foreach ($cacheVars as $var => $data) {
            $current = $data['current'];
            $default = $data['default'];
            $defaultInfo = $default !== null ? " (default: $default)" : " (no default)";
            echo "  $var: $current$defaultInfo\n";
        }
    } else {
        echo "  No cache-related variables found.\n";
    }
    
    echo "\n";
    
    // Count variables by category
    echo "Variable Statistics:\n";
    echo "  Total Variables: " . count($mysql_variables) . "\n";
    
    // Count variables starting with common prefixes
    $prefixes = ['innodb_', 'myisam_', 'query_cache_', 'thread_', 'log_'];
    foreach ($prefixes as $prefix) {
        $count = 0;
        foreach ($mysql_variables as $var => $data) {
            if (strpos($var, $prefix) === 0) {
                $count++;
            }
        }
        if ($count > 0) {
            echo "  Variables starting with '$prefix': $count\n";
        }
    }
    
} else {
    echo "Error: mysql_variables.php file not found!\n";
    echo "Please run extract_mysql_variables.php first to generate the variables file.\n";
}
?> 