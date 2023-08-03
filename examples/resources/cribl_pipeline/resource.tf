resource "cribl_pipeline" "my_pipeline" {
    conf = {
        async_func_timeout = 6
        description = "...my_description..."
        functions = [
            {
                conf = {}
                description = "...my_description..."
                disabled = false
                filter = "...my_filter..."
                final = true
                group_id = "...my_groupId..."
                id = "d9d8d69a-674e-40f4-a7cc-8796ed151a05"
            },
        ]
        groups = {
            "repellendus" = {
                description = "...my_description..."
                disabled = true
                name = "Fred Strosin"
            }
            "molestiae" = {
                description = "...my_description..."
                disabled = false
                name = "Erik Lebsack"
            }
        }
        output = "...my_output..."
        streamtags = [
            "...",
        ]
    }
            id = "1ba928fc-8167-442c-b739-205929396fea"
        }