-- Rename default alert rules from Chinese to English.
-- Uses UPDATE ... WHERE name = to only affect the original seeded rules.

UPDATE ops_alert_rules SET name = 'High Error Rate',           description = 'Alert when error rate exceeds 5% for 5 minutes',                                updated_at = NOW() WHERE name = '错误率过高';
UPDATE ops_alert_rules SET name = 'Low Success Rate',          description = 'Alert when success rate drops below 95% for 5 minutes (service availability degraded)', updated_at = NOW() WHERE name = '成功率过低';
UPDATE ops_alert_rules SET name = 'P99 Latency Too High',      description = 'Alert when P99 latency exceeds 3000ms for 10 minutes',                          updated_at = NOW() WHERE name = 'P99延迟过高';
UPDATE ops_alert_rules SET name = 'P95 Latency Too High',      description = 'Alert when P95 latency exceeds 2000ms for 10 minutes',                          updated_at = NOW() WHERE name = 'P95延迟过高';
UPDATE ops_alert_rules SET name = 'CPU Usage Too High',        description = 'Alert when CPU usage exceeds 85% for 10 minutes',                               updated_at = NOW() WHERE name = 'CPU使用率过高';
UPDATE ops_alert_rules SET name = 'Memory Usage Too High',     description = 'Alert when memory usage exceeds 90% for 10 minutes (potential OOM risk)',        updated_at = NOW() WHERE name = '内存使用率过高';
UPDATE ops_alert_rules SET name = 'Concurrency Queue Buildup', description = 'Alert when queue depth exceeds 100 for 5 minutes (insufficient processing capacity)', updated_at = NOW() WHERE name = '并发队列积压';
UPDATE ops_alert_rules SET name = 'Critical Error Rate',       description = 'Alert when error rate exceeds 20% for 1 minute (severe service failure)',        updated_at = NOW() WHERE name = '错误率极高';
