# Data analysis

## Setup

Run the following to populate the DB locally

> Ensure to run `make migrate-up` in project root to initialise tables and relations

```sh
poetry install \
source .venv/bin/activate \
python init_dataset.py \
python populate_db.py
```

## Credits

[A multilevel carbon and water footprint dataset of food commodities](https://doi.org/10.1038/s41597-021-00909-8)
