# file-manager
File manager

- Follow the API spec https://platform.openai.com/docs/api-reference/files. We can start from using gRPC if it works.
- Also consider Hugging Face Dataset API (https://huggingface.co/docs/hub/en/datasets-overview)
- Use SQL for metadata management
- Use https://min.io/ as object store, but it would be great if we can replace
- Support the `standalone` mode for easy debug (https://github.com/llm-operator/job-manager/blob/main/dispatcher/internal/config/config.go#L15)
- Include `TenantID` in each file record so that we can later implement tenant isolation.
- Use OpenAI python client to verify the implementation.
