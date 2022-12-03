"""Clean SuEatableLife dataset to populate Postgres database

1. Merge columns between two sheets
2. Filter to keep only relevant columns
3. Simplify column names
4. Update product names to remove product suffixes
"""

import enum

import pandas as pd
import numpy as np
from numpy import NaN

REGIONS = ["Africa", "America", "Asia", "Europe", "Mediterranean", "Oceania", "World"]


class Greenhouse(enum.Enum):
    NA = 1
    NOT_HEATED = 2
    HEATED = 3


df1 = pd.read_excel(
    "SuEatableLife_Food_Fooprint_database.xlsx", sheet_name="SEL CF DATA SOURCES"
)
df2 = pd.read_excel(
    "SuEatableLife_Food_Fooprint_database.xlsx",
    sheet_name="SEL CF for users",
)

df2.rename(
    columns={
        "Food commodity TYPOLOGY": "FOOD COMMODITY TYPOLOGY",
        "Food commodity ITEM": "FOOD COMMODITY ITEM",
    },
    inplace=True,
)

df1["FOOD COMMODITY ITEM"] = df1["FOOD COMMODITY ITEM"].replace(
    {
        "BEER GLASS BOTTLE": "BEER IN GLASS",
        "TEMPÈ": "TEMPEH",
        "STRAWBERRY & BANANA JUICE (I)": "STRAWBERRY JUICE (I)",
        "DURUM WHEAT SEMOLINA": "DURUM WHEAT",
        "OATMEAL FLOUR": "OAT MEAL",
        "WHEAT FLOUR": "WHEAT PLAIN FLOUR",
        "CHOCOLATE (F)": "CHOCOLATE",
        "ICE CREAM (F)": "ICE CREAM",
        "VANILLA (F)": "VANILLA",
        "MUNG BEAN": "MUNG BEAN FLOUR",
        "CAMEMBERT CHEESE": "CAMEMBERT",
        "PARMIGGIANO REGGIANO": "PARMIGIANO REGGIANO",
        "COW'S MILK": "COW MILK",
        "DUCK BONE FREE MEAT": "DUCK",
        "TURKEY BONE FREE MEAT": "TURKEY",
        "COD (FISH STICK)": "COD FISH STICK",
        "HAKE (FISH STICK)": "HAKE FISH STICK",
        "MACKERAL (FISH STICK)": "MACKEREL FISH STICK",
        "POLLOCK (FISH STICK) (F)": "POLLOCK STICK (F)",
        "KIWI FRUIT": "KIWI",
        "EGGPLANT (AUBERGIN)": "EGGPLANT",
        "EGGPLANT (AUBERGIN) (g)": "EGGPLANT (g)",
        "SWEDE (RUTABAGE)": "SWEDE",
        "TURNIPS": "TURNIP",
        "CARROT (F)": "CARROTS (F)",
        "BEAN (F)": "BEANS (F)",
        "GREEN BEAN (F)": "GREEN BEANS (F)",
        "PRAWNS/SHRIMP": "PRAWNS/SHRIMPS",
        "PRAWNS/SHRIMP (F)": "PRAWNS/SHRIMPS (F)",
        "PORK MEAT WITH BONE": "PORK WITH BONE",
    }
)

df1["FOOD COMMODITY TYPOLOGY"] = df1["FOOD COMMODITY TYPOLOGY"].replace(
    {
        "FRUIT JUICE (imported ingredients)": "FRUIT JUICE IMPORTED",
        "VEGETABLES GREENHOUSE NOT HEATED": "VEGETABLES NOT HEATED GREENHOUSE",
    }
)
df2["FOOD COMMODITY TYPOLOGY"] = df2["FOOD COMMODITY TYPOLOGY"].replace(
    {
        "FRUIT JUICE (imported ingredients)": "FRUIT JUICE IMPORTED",
        "TOMATO PUREE,PEELED, CHOPPED": "TOMATO PUREE, PEELED, CHOPPED",
        "CHEESE FRESH": "FRESH CHEESE",
        "CHEESE HARD & SEMI-HARD": "HARD AND SEMI-HARD CHEESE",
    }
)
df2["FOOD COMMODITY ITEM"] = df2["FOOD COMMODITY ITEM"].map(lambda x: x.rstrip("*"))
df2["FOOD COMMODITY ITEM"] = df2["FOOD COMMODITY ITEM"].replace(
    {
        "BREAD FROZEN (F)": "ROLLS (F)",
        "COFFEE DRIP FILTERED (L)": "COFFEE DRIPPED FILTERED (L)",
        "COFFEE SOLUBLE POWDER (L)": "COFFEE SOLUBLE (L)",
        "TEMPE'": "TEMPEH",
        "QUORNE": "QUORN",
        "CHEESE SEMI-HARD": "SEMI-HARD CHEESE",
        "CEDDAR": "CHEDDAR",
        "NECK": "PORK NECK",
        "PORK SAUSAGES": "PORK SAUSAGE",
        "DUCK MEAT BONE FREE": "DUCK",
        "TURKEY MEAT BONE FREE": "TURKEY",
        "CABAGGE": "CABBAGE",
        "MAKEREL FISH STICK": "MACKEREL FISH STICK",
        "RAPE SEED": "RAPESEED",
        "SESAM SEED": "SESAME SEED",
        "SWARDFISH": "SWORDFISH",
        "CHILLY": "CHILLI",
        "GERKIN": "GHERKIN",
        "GERKIN (G)": "GHERKIN (G)",
        "RASPBERRIES (F)": "RASPBERRY (F)",
        "PORK MEAT WITH BONE": "PORK WITH BONE",
    }
)
df2["FOOD COMMODITY TYPOLOGY"] = (
    df2["FOOD COMMODITY TYPOLOGY"].str.replace(r"\(.*?\)", "", regex=True).str.strip()
)
df2["FOOD COMMODITY TYPOLOGY"] = df2["FOOD COMMODITY TYPOLOGY"].map(
    lambda x: x.rstrip("*")
)
df2["FOOD COMMODITY ITEM"] = df2["FOOD COMMODITY ITEM"].map(
    lambda x: x.replace(" (fresh)", "")
)

df1["FOOD COMMODITY ITEM"] = df1["FOOD COMMODITY ITEM"].str.strip()
df1["FOOD COMMODITY TYPOLOGY"] = df1["FOOD COMMODITY TYPOLOGY"].str.strip()
df2["FOOD COMMODITY ITEM"] = df2["FOOD COMMODITY ITEM"].str.strip()


# print(df2["FOOD COMMODITY ITEM"].value_counts().to_string())
# print(df2["FOOD COMMODITY TYPOLOGY"].value_counts().to_string())

new_df = pd.merge(
    df1,
    df2,
    how="left",
    right_on=["FOOD COMMODITY ITEM", "FOOD COMMODITY TYPOLOGY"],
    left_on=["FOOD COMMODITY ITEM", "FOOD COMMODITY TYPOLOGY"],
)

print(len(new_df))
print(
    new_df.loc[
        new_df["Suggested CF value"].isnull(),
        ["FOOD COMMODITY ITEM", "FOOD COMMODITY TYPOLOGY"],
    ]
    .drop_duplicates()
    .to_string(index=False)
)

new_df.dropna(subset=["Suggested CF value"], inplace=True)


new_df = new_df.filter(
    [
        "FOOD COMMODITY TYPOLOGY",
        "FOOD COMMODITY ITEM",
        "Food commodity sub-TYPOLOGY",
        "Carbon footprint  (kg CO2 eq/kg or litre of food commodity)",
        "Type of source",
        "Full Reference",
        "Publication year",
        "Food processing",
        "Region",
        "Suggested CF value",
    ]
)


new_df.rename(
    columns={
        "FOOD COMMODITY TYPOLOGY": "Typology",
        "Food commodity sub-TYPOLOGY": "Sub Typology",
        "FOOD COMMODITY ITEM": "Name",
        "Carbon footprint  (kg CO2 eq/kg or litre of food commodity)": "Carbon footprint",
        "Suggested CF value": "Suggested CF",
    },
    inplace=True,
)

# new_df["Imported"] = False
# new_df["Frozen"] = False
# new_df["Category"] = NaN
# new_df["Greenhouse"] = Greenhouse.NA.value


def update_label_with_format(x: str, remove: str, format: str):
    x = x.replace(remove, "")
    return format.format(x)


def update_frozen_items():
    frozen = [
        "SHELLFISH FROZEN",
        "FRUIT FROZEN",
        "LEGUMES FROZEN",
        "VEGETABLES FROZEN",
        "FISH FROZEN",
        "BREAD FROZEN",
        "ICE CREAM",
    ]

    for typology in frozen:
        indices = new_df.index[new_df["Typology"] == typology]
        df = new_df.loc[indices]
        df["Name"] = df["Name"].apply(update_label_with_format, args=(" (F)", "{}"))
        # df["Frozen"] = True
        new_df.loc[indices] = df


update_frozen_items()


def update_imported_items():
    imported = ["FRUIT JUICE IMPORTED", "FRUIT IMPORTED"]
    for typology in imported:
        indices = new_df.index[new_df["Typology"] == typology]
        df = new_df.loc[indices]
        df["Name"] = df["Name"].apply(update_label_with_format, args=(" (I)", "{}"))
        # df["Imported"] = True
        new_df.loc[indices] = df


update_imported_items()


def update_greenhouse_not_heated():
    typologies = [
        "FRUIT GREENHOUSE NOT HEATED",
        "VEGETABLES NOT HEATED GREENHOUSE",
        "LEGUMES NOT HEATED GREENHOUSE",
    ]
    for typology in typologies:
        indices = new_df.index[new_df["Typology"] == typology]
        df = new_df.loc[indices]
        df["Name"] = df["Name"].apply(update_label_with_format, args=(" (g)", "{}"))
        new_df.loc[indices] = df


update_greenhouse_not_heated()


def update_greenhouse_heated():
    typologies = ["VEGETABLES HEATED GREENHOUSE", "FRUIT HEATED GREENHOUSE"]
    for typology in typologies:
        indices = new_df.index[new_df["Typology"] == typology]
        df = new_df.loc[indices]
        df["Name"] = df["Name"].apply(update_label_with_format, args=(" (G)", "{}"))
        new_df.loc[indices] = df


update_greenhouse_heated()

def update_liquid_items():
    typologies = ["COFFEE LIQUID"]
    for typology in typologies:
        indices = new_df.index[new_df["Typology"] == typology]
        df = new_df.loc[indices]
        df["Name"] = df["Name"].apply(update_label_with_format, args=(" (L)", "{}"))
        new_df.loc[indices] = df

update_liquid_items()

def map_region_name(old_name: str) -> str:
    map = {
        "Africa": "Africa",
        "Asia": "Asia",
        "Atlantic Ocean": "World",
        "C America": "America",
        "C Europe": "Europe",
        "N Europe": "Europe",
        "CS America": "America",
        "E Europe": "Europe",
        "Europe": "Europe",
        "Mediterranean area": "Mediterranean",
        "N Africa": "Africa",
        "N America": "America",
        "non UE countries": "Europe",
        "Oceania": "Oceania",
        "Other imported": "World",
        "Pacific Ocean": "World",
        "S Africa": "Africa",
        "S America": "America",
        "W America": "America",
        "W Europe": "Europe",
        "World": "World",
    }

    try:
        if (new_name := map[old_name]) in REGIONS:
            return new_name
        else:
            print("Mapped region not in REGIONS const")
            raise KeyError()
    except KeyError:
        return "World"


new_df["Region"] = new_df["Region"].apply(map_region_name)

new_df["Sub Typology"] = np.where(
    new_df["Sub Typology"] == "-", NaN, new_df["Sub Typology"]
)

# df1["Typology"] = df1["Typology"].replace(
#     {
#         "FRUIT OPENFIELD": "FRUIT",
#         "FRUIT IMPORTED": "FRUIT",
#         "VEGETABLES OPENFIELD": "VEGETABLES",
#         # "VEGETABLES HEATED GREENHOUSE": "VEGETABLES",
#         # "VEGETABLES GREENHOUSE NOT HEATED": "VEGETABLES",
#         "BEEF BONE FREE MEAT": "BEEF",
#         "BEEF MEAT WITH BONE": "BEEF",
#         "PORK BONE FREE MEAT": "PORK",
#         "PORK MEAT WITH BONE": "PORK",
#         "POULTRY BONE FREE MEAT": "POULTRY",
#         "POULTRY MEAT WITH BONE": "POULTRY",
#         "LAMB BONE FREE MEAT": "LAMB",
#         "LAMB MEAT WITH BONE": "LAMB",
#         "BUFFALO BONE FREE MEAT": "BUFFALO",
#         "STARCHY TUBERS": "VEGETABLES",
#         "FRUIT JUICE LOCAL": "FRUIT JUICE",
#         "FRUIT JUICE (imported ingredients)": "FRUIT JUICE",
#         "TOMATO PUREE, PEELED, CHOPPED": "SAUCE",
#         "TOMATO SAUCE": "SAUCE",
#         "FISH PROCESSED": "FISH",
#         "SHELLFISH FROZEN": "FISH",
#         "FRUIT FROZEN": "FRUIT",
#         # "FRUIT GREENHOUSE NOT HEATED": "FRUIT",
#         # "FRUIT HEATED GREENHOUSE": "FRUIT",
#         # "LEGUMES NOT HEATED GREENHOUSE": "LEGUMES",
#         "LEGUMES FROZEN": "LEGUMES",
#         "VEGETABLES FROZEN": "VEGETABLES",
#         "COFFEE GREEN": "COFFEE",
#         "COFFEE LIQUID": "COFFEE",
#         "COFFEE GROUND & PARCHMENT": "COFFEE",
#         "VEGETAL MILK": "PLANT MILK",
#         "YEAST LIQUID": "YEAST",
#         "YEAST COMPRESSED": "YEAST",
#         "YEAST DRIED": "YEAST",
#         "EGG PASTA": "PASTA",
#         "VEGETABLES CANNED": "VEGETABLES",
#         "LEGUMES CANNED": "LEGUMES",
#         "RABBIT MEAT WITH BONE": "RABBIT",
#         "NUTS COVERED WITH CHOCOLATE": "NUTS",
#         "CRISPBREAD": "CRACKERS",
#         "MUSHROOM": "VEGETABLES",
#         "NUTS PASTE": "NUTS",
#         "EDIBLE INSECTS": "INSECTS",
#         "SNAILS": "INSECTS",
#         "PESTO": "SAUCE",
#         "GRAIN FLOUR": "FLOUR",
#         "LEGUME FLOUR": "FLOUR",
#         "PORK CURED MEAT": "PORK",
#         "BREAKFAST CEREALS": "GRAINS",
#         "KANGAROO MEAT": "KANGAROO",
#         "FISH FROZEN": "FISH",
#         "FRESH CHEESE": "CHEESE",
#         "HARD AND SEMI-HARD CHEESE": "CHEESE",
#         "BREAD FROZEN": "BREAD",
#     }
# )

new_df["Sub Typology"] = new_df["Sub Typology"].str.replace(
    " (SUB TYP)", "", regex=False
)
new_df["Sub Typology"] = new_df["Sub Typology"].str.replace(
    " (SUB-TYP)", "", regex=False
)
new_df["Sub Typology"] = new_df["Sub Typology"].str.replace("*", "", regex=False)

new_df["Suggested CF"] = new_df["Suggested CF"].replace(
    {
        "OK item": "ITEM",
        "BETTER TYPOLOGY": "TYPOLOGY",
        "Item matching typology": "ITEM",
        "better typol. or subtypol.": "TYPOLOGY",
        "better typology": "TYPOLOGY",
        "item or typology": "ITEM",
        "better subtypology": "SUBTYPOLOGY",
        "item or typology or subtypology": "SUBTYPOLOGY",
    }
)


new_df.to_excel("data.xlsx")
