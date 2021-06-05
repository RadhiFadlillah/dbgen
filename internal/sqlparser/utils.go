package sqlparser

// extractQueryParams extract parameters that used in a query, also possible arguments to fill the query.
func extractQueryParams(sqlQuery string) (string, []Parameter, map[string]interface{}) {
	// Extract unique parameter names from query
	var paramNames []string
	queryArgs := make(map[string]interface{})
	mapRequiredParam := make(map[string]bool)

	sqlQuery = rxQueryParams.ReplaceAllStringFunc(sqlQuery, func(s string) string {
		parts := rxQueryParams.FindStringSubmatch(s)

		paramName := parts[2]
		paramRequired := parts[1] == "!"
		if _, exist := mapRequiredParam[paramName]; !exist {
			paramNames = append(paramNames, paramName)
		}

		queryArgs[paramName] = 1
		mapRequiredParam[paramName] = mapRequiredParam[paramName] || paramRequired
		return ":" + paramName
	})

	// Convert param names to list of parameter
	var params []Parameter
	for _, paramName := range paramNames {
		params = append(params, Parameter{
			Name:     paramName,
			Required: mapRequiredParam[paramName],
		})
	}

	return sqlQuery, params, queryArgs
}
