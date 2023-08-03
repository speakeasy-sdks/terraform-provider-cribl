data "cribl_pipeline" "my_pipeline" {
    conf = {
        async_func_timeout = 4
        description = "...my_description..."
        functions = [
            {
                conf = {}
                description = "...my_description..."
                disabled = true
                filter = "...my_filter..."
                final = true
                group_id = "...my_groupId..."
                id = "6eb10faa-a235-42c5-9559-07aff1a3a2fa"
            },
        ]
        groups = {
            "occaecati" = {
                description = "...my_description..."
                disabled = false
                name = "Claudia Krajcik"
            }
            "quia" = {
                description = "...my_description..."
                disabled = false
                name = "Kayla O'Kon"
            }
        }
        output = "...my_output..."
        streamtags = [
            "...",
        ]
    }
            id = "c3f5ad01-9da1-4ffe-b8f0-97b0074f1547"
        }