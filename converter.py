import os
import re
import csv
import shutil

TABLE_ROWS_PATTERN = r'([^\n]*?\|[^\n]*?\n)'
TABLE_AREA_PATTERN = r'^\#\#\#\#\#\# ([a-zA-Z0-9 ]+)\n*(?:\n?[^\n]*?\|[^\n]*?\n)+'

CSV_FOLDER = "./csv/"
STRUCT_FOLDER = CSV_FOLDER + "structs/"
CONSTS_FOLDER = CSV_FOLDER + "consts/"


def create_filename(title):
    return title.title().replace(" ", "")

def is_const(title):
    return title.endswith(("Level", "Tier", "Flags", "Types"))

# cleanup
shutil.rmtree(CSV_FOLDER, ignore_errors=True)
os.mkdir(CSV_FOLDER)
os.mkdir(STRUCT_FOLDER)
os.mkdir(CONSTS_FOLDER)

def create_csv_files(textfile):
    filetext = textfile.read()
    textfile.close()
    matches = re.finditer(TABLE_AREA_PATTERN, filetext, re.MULTILINE)

    for matchNum, match in enumerate(matches, start=1):
        title = match.group(1)
        if title.endswith("JSON Params") or title.endswith("Query String Params"):
            continue

        rows = re.findall(TABLE_ROWS_PATTERN, match.group())

        cells = [[]]
        for row in rows:
            a = "".join(set(row.strip().replace("|", "-"))).strip()
            if a == "-":
                continue

            cols = [x.strip().replace("\\*", "*") for x in row.strip().split('|')]
            cells.append(list(filter(None, cols)))
        if len(cells[0]) == 0:
            cells = cells[1:]

        # make sure each row has the same width
        width = 0
        for row in cells:
            if len(row) > width:
                width = len(row)
        for i, row in enumerate(cells):
            for x in range(len(row), width):
                cells[i].append("")

        path = ""
        if title.__contains__("Structure"):
            path = STRUCT_FOLDER
        elif is_const(title):
            path = CONSTS_FOLDER

        if path == "":
            print("--" + title)
            continue

        with open(path + create_filename(title) + ".csv", mode='w') as f:
            writer = csv.writer(f, delimiter=',', quotechar='"', quoting=csv.QUOTE_MINIMAL)
            for row in cells:
                writer.writerow(row)


def iterate_markdown_files(path):
    for file in os.listdir(path):
        if file.endswith(".md"):
            f = open(os.path.join(path, file), "r")
            create_csv_files(f)

BASE = "./discord-api-docs/docs/"
iterate_markdown_files(BASE + "resources/")
iterate_markdown_files(BASE + "topics/")