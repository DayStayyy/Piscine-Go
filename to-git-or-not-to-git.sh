#!/bin/bash
message=$(curl -s https://ytrack.learn.ynov.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"ACLAMADIEU\"}}){id}}"}')
prefix='{"data":{"user":[{"id":';
suffix="}]}}";
foo=${message#"$prefix"}
foo=${foo%"$suffix"}
echo "${foo}"