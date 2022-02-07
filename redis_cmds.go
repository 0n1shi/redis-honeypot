package main

import "strings"

const (
	redisNewLine    = "\r\n"
	redisNil        = "$-1" + redisNewLine
	redisMsgOK      = "+OK" + redisNewLine
	redisMsgSuccess = ":1" + redisNewLine
	redisMsgFailure = ":0" + redisNewLine
)

var redisDataMap = map[string]string{}

func redisCOMMAND() string {
	return redisMsgOK
}

func redisPING() string {
	return "+PONG" + redisNewLine
}

func redisKEYS() string { // TODO: return keys
	keys := []string{}
	for k := range redisDataMap {
		keys = append(keys, k)
	}
	return toRedisStrArray(keys)
}

func redisSET(key_value []string) string {
	redisDataMap[key_value[0]] = key_value[1]
	return redisMsgOK
}

func redisGET(key string) string {
	if val, ok := redisDataMap[key]; ok {
		return toRedisStr(val)
	}
	return redisNil
}

func redisDEL(key string) string {
	if _, ok := redisDataMap[key]; !ok {
		return redisMsgFailure
	}
	delete(redisDataMap, key)
	return redisMsgSuccess
}

func redisINFO() string {
	msgs := []string{
		"# Server",
		"redis_version:6.2.6",
		"redis_git_sha1:00000000",
		"redis_git_dirty:0",
		"redis_build_id:b61f37314a089f19",
		"redis_mode:standalone",
		"os:Linux 5.10.76-linuxkit x86_64",
		"arch_bits:64",
		"multiplexing_api:epoll",
		"atomicvar_api:atomic-builtin",
		"gcc_version:10.2.1",
		"process_id:1",
		"process_supervised:no",
		"run_id:b7e246c00e755f15698fa5fb55e38b241db258f3",
		"tcp_port:6379",
		"server_time_usec:1644233854325059",
		"uptime_in_seconds:103028",
		"uptime_in_days:1",
		"hz:10",
		"configured_hz:10",
		"lru_clock:66686",
		"executable:/data/redis-server",
		"config_file:",
		"io_threads_active:0",
		"",
		"# Clients",
		"connected_clients:1",
		"cluster_connections:0",
		"maxclients:10000",
		"client_recent_max_input_buffer:24",
		"client_recent_max_output_buffer:0",
		"blocked_clients:0",
		"tracking_clients:0",
		"clients_in_timeout_table:0",
		"",
		"# Memory",
		"used_memory:874664",
		"used_memory_human:854.16K",
		"used_memory_rss:7929856",
		"used_memory_rss_human:7.56M",
		"used_memory_peak:932680",
		"used_memory_peak_human:910.82K",
		"used_memory_peak_perc:93.78%",
		"used_memory_overhead:830456",
		"used_memory_startup:809880",
		"used_memory_dataset:44208",
		"used_memory_dataset_perc:68.24%",
		"allocator_allocated:1090640",
		"allocator_active:1359872",
		"allocator_resident:3874816",
		"total_system_memory:2082197504",
		"total_system_memory_human:1.94G",
		"used_memory_lua:37888",
		"used_memory_lua_human:37.00K",
		"used_memory_scripts:0",
		"used_memory_scripts_human:0B",
		"number_of_cached_scripts:0",
		"maxmemory:0",
		"maxmemory_human:0B",
		"maxmemory_policy:noeviction",
		"allocator_frag_ratio:1.25",
		"allocator_frag_bytes:269232",
		"allocator_rss_ratio:2.85",
		"allocator_rss_bytes:2514944",
		"rss_overhead_ratio:2.05",
		"rss_overhead_bytes:4055040",
		"mem_fragmentation_ratio:9.53",
		"mem_fragmentation_bytes:7097952",
		"mem_not_counted_for_evict:0",
		"mem_replication_backlog:0",
		"mem_clients_slaves:0",
		"mem_clients_normal:20504",
		"mem_aof_buffer:0",
		"mem_allocator:jemalloc-5.1.0",
		"active_defrag_running:0",
		"lazyfree_pending_objects:0",
		"lazyfreed_objects:0",
		"",
		"# Persistence",
		"loading:0",
		"current_cow_size:0",
		"current_cow_size_age:0",
		"current_fork_perc:0.00",
		"current_save_keys_processed:0",
		"current_save_keys_total:0",
		"rdb_changes_since_last_save:0",
		"rdb_bgsave_in_progress:0",
		"rdb_last_save_time:1644233798",
		"rdb_last_bgsave_status:ok",
		"rdb_last_bgsave_time_sec:0",
		"rdb_current_bgsave_time_sec:-1",
		"rdb_last_cow_size:192512",
		"aof_enabled:0",
		"aof_rewrite_in_progress:0",
		"aof_rewrite_scheduled:0",
		"aof_last_rewrite_time_sec:-1",
		"aof_current_rewrite_time_sec:-1",
		"aof_last_bgrewrite_status:ok",
		"aof_last_write_status:ok",
		"aof_last_cow_size:0",
		"module_fork_in_progress:0",
		"module_fork_last_cow_size:0",
		"",
		"# Stats",
		"total_connections_received:3",
		"total_commands_processed:11",
		"instantaneous_ops_per_sec:0",
		"total_net_input_bytes:263",
		"total_net_output_bytes:69430",
		"instantaneous_input_kbps:0.00",
		"instantaneous_output_kbps:0.00",
		"rejected_connections:0",
		"sync_full:0",
		"sync_partial_ok:0",
		"sync_partial_err:0",
		"expired_keys:0",
		"expired_stale_perc:0.00",
		"expired_time_cap_reached_count:0",
		"expire_cycle_cpu_milliseconds:9378",
		"evicted_keys:0",
		"keyspace_hits:1",
		"keyspace_misses:1",
		"pubsub_channels:0",
		"pubsub_patterns:0",
		"latest_fork_usec:12031",
		"total_forks:1",
		"migrate_cached_sockets:0",
		"slave_expires_tracked_keys:0",
		"active_defrag_hits:0",
		"active_defrag_misses:0",
		"active_defrag_key_hits:0",
		"active_defrag_key_misses:0",
		"tracking_total_keys:0",
		"tracking_total_items:0",
		"tracking_total_prefixes:0",
		"unexpected_error_replies:0",
		"total_error_replies:2",
		"dump_payload_sanitizations:0",
		"total_reads_processed:14",
		"total_writes_processed:11",
		"io_threaded_reads_processed:0",
		"io_threaded_writes_processed:0",
		"",
		"# Replication",
		"role:master",
		"connected_slaves:0",
		"master_failover_state:no-failover",
		"master_replid:64e63f69089ed3cc9e898938e619f3f28b5d3f6d",
		"master_replid2:0000000000000000000000000000000000000000",
		"master_repl_offset:0",
		"second_repl_offset:-1",
		"repl_backlog_active:0",
		"repl_backlog_size:1048576",
		"repl_backlog_first_byte_offset:0",
		"repl_backlog_histlen:0",
		"",
		"# CPU",
		"used_cpu_sys:181.982214",
		"used_cpu_user:33.850671",
		"used_cpu_sys_children:0.014989",
		"used_cpu_user_children:0.001187",
		"used_cpu_sys_main_thread:181.965005",
		"used_cpu_user_main_thread:33.842315",
		"",
		"# Modules",
		"",
		"# Errorstats",
		"errorstat_ERR:count=1",
		"errorstat_WRONGPASS:count=1",
		"",
		"# Cluster",
		"cluster_enabled:0",
		"",
		"# Keyspace",
		"db0:keys=1,expires=0,avg_ttl=0",
		"",
	}
	return toRedisStr(strings.Join(msgs, "\r\n"))
}
