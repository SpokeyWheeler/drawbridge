---
roles:
  - readonly:
    databases:
    - test
      schemas:
        - public
          tables:
            - foo
    permissions:
      - select
