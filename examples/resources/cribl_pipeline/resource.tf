resource "cribl_pipeline" "my_pipeline" {
    conf = {
        async_func_timeout = 7
        description = "...my_description..."
        functions = [
            {
                conf = {}
                description = "...my_description..."
                disabled = true
                filter = "...my_filter..."
                final = false
                group_id = "...my_groupId..."
                id = "dfc2ddf7-cc78-4ca1-ba92-8fc816742cb7"
            },
        ]
        groups = {
            "ipsum" = {
                description = "...my_description..."
                disabled = true
                name = "Dorothy Hane"
            }
            "iste" = {
                description = "...my_description..."
                disabled = false
                name = "Lester Welch"
            }
        }
        output = "...my_output..."
        streamtags = [
            "...",
        ]
    }
            id = "7596eb10-faaa-4235-ac59-55907aff1a3a"
        }