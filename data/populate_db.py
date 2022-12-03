import collections
import requests
import pandas as pd

REGIONS = ["Africa", "America", "Asia", "Europe", "Mediterranean", "Oceania", "World"]
SERVICE = "http://localhost:8080/api/neutral_diet.food.v1.FoodService"

CF_TYPE_MAP = {"TYPOLOGY": 1, "SUBTYPOLOGY": 2, "ITEM": 3}

for region in REGIONS:
    requests.post(f"{SERVICE}/CreateRegion", json={"region": {"name": region}})

sub_typology_index: dict[str, int] = {}
typology_index: dict[str, int] = {}
typology_and_sub_typology_index: dict[str, dict[str, int]] = collections.defaultdict(
    dict
)
source_index: dict[str, int] = {}
food_name_index: dict[str, int] = {}

df = pd.read_excel("data.xlsx")

sub_typologies: list[str] = df["Sub Typology"].drop_duplicates().dropna().to_list()

id = 1
for sub_typology in sub_typologies:
    res = requests.post(
        f"{SERVICE}/CreateSubTypology", json={"sub_typology": {"name": sub_typology}}
    )
    if not res.ok:
        res.raise_for_status()
    id = int(res.json()["id"])
    sub_typology_index[sub_typology] = id

for index, row in df.iterrows():
    print(f"Adding {row['Name']}")

    source_id: int
    if row["Full Reference"] not in source_index:
        source_response = requests.post(
            f"{SERVICE}/CreateSource",
            json={
                "source": {
                    "reference": row["Full Reference"],
                    "year": int(row["Publication year"]),
                    "region_name": row["Region"],
                }
            },
        )
        if not source_response.ok:
            source_response.raise_for_status()
        source_id = int(source_response.json()["id"])
        source_index[row["Full Reference"]] = source_id
    else:
        source_id = source_index[row["Full Reference"]]

    typology_id: int
    sub_typology_id = sub_typology_index.get(row["Sub Typology"])
    if sub_typology_id is not None:
        if row["Sub Typology"] not in typology_and_sub_typology_index.get(
            row["Typology"], {}
        ):
            typology_response = requests.post(
                f"{SERVICE}/CreateTypology",
                json={
                    "typology": {
                        "name": row["Typology"],
                        "sub_typology_id": sub_typology_id,
                    }
                },
            )
            if not typology_response.ok:
                typology_response.raise_for_status()
            typology_id = int(typology_response.json()["id"])
            typology_and_sub_typology_index[row["Typology"]][
                row["Sub Typology"]
            ] = typology_id
        else:
            typology_id = typology_and_sub_typology_index[row["Typology"]][
                row["Sub Typology"]
            ]
    else:
        if row["Typology"] not in typology_index:
            typology_response = requests.post(
                f"{SERVICE}/CreateTypology",
                json={"typology": {"name": row["Typology"]}},
            )
            if not typology_response.ok:
                typology_response.raise_for_status()
            typology_id = int(typology_response.json()["id"])
            typology_index[row["Typology"]] = typology_id
        else:
            typology_id = typology_index[row["Typology"]]

    if row["Name"] not in food_name_index:
        res = requests.post(
            f"{SERVICE}/CreateFoodItem",
            json={
                "food_item": {
                    "name": row["Name"],
                    "typology_id": typology_id,
                    "cf_type": CF_TYPE_MAP[row["Suggested CF"]],
                }
            },
        )
        if not res.ok:
            res.raise_for_status()

        food_item_id = int(res.json()["id"])
        food_name_index[row["Name"]] = food_item_id
    else:
        food_item_id = food_name_index[row["Name"]]

    requests.post(
        f"{SERVICE}/CreateLifeCycle",
        json={
            "life_cycle": {
                "carbon_footprint": row["Carbon footprint"],
                "food_item_id": food_item_id,
                "source_id": source_id,
            }
        },
    )
