data "Cribl_pipeline" "my_pipeline" {
    conf = {
        async_func_timeout = 3
        description = "...my_description..."
        functions = [
            {
                conf = {}
                description = "...my_description..."
                disabled = false
                filter = "...my_filter..."
                final = true
                group_id = "...my_group_id..."
                id = "3c969e9a-3efa-477d-bb14-cd66ae395efb"
            },
        ]
        groups = {
            "provident" = {
                description = "...my_description..."
                disabled = false
                name = "Nelson Lesch"
            }
            "deserunt" = {
                description = "...my_description..."
                disabled = false
                name = "Ada Moen IV"
            }
        }
        output = "...my_output..."
        streamtags = [
            "...",
        ]
    }
            id = "4ba4469b-6e21-4419-9989-0afa563e2516"
        }