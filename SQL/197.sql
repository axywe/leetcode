SELECT w2.id FROM Weather w1
JOIN Weather w2 ON w2.recordDate = w1.recordDate+1
WHERE w1.temperature < w2.temperature