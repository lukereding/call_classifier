import os
from os.path import isfile, join
from sys import exit
from typing import Dict, List

import numpy as np

import cv2


def defineKeyboardMapping() -> Dict[int, str]:
    """
    Defines a mapping between keyboard codes and what key you
    actually pressed.
    """

    mapping = {
        48: "0",
        49: "1",
        50: "2",
        51: "3",
        52: "4",
        53: "5",
        54: "6",
        55: "7",
        56: "8",
        57: "9",
        119: "w",
        99: "c",
        32: " ",
        117: "u",
    }
    return mapping


def defineKeyToCategory() -> Dict[str, str]:
    """
    Defines the mapping between keys and categories.
    """
    return {
        "0": "whine",
        "1": "whine_one_chuck",
        "2": "whine_two_chucks",
        "3": "whine_three_chucks",
        "4": "whine_four_chucks",
        "5": "whine_five_chucks",
        "6": "whine_six_chucks",
        "7": "whine_seven_chucks",
        "8": "whine_eight_chucks",
        "9": "mew",
        "u": "unknown",
    }


def getFolder() -> str:
    """
    Get the current working directory.
    """
    return os.getcwd()


def getListOfPNGsIn(dir: str) -> List[str]:
    """
    Return the full path of files in `dir`.
    """

    onlyfiles = [f for f in os.listdir(dir) if isfile(join(dir, f))]
    pngs = [file for file in onlyfiles if file.endswith(".png")]
    return pngs


def showImageToUser(filename: str) -> str:
    """
    Read in the png defined in `filename` and show it to the user.
    The user presses a key and the category associated with the key is returned.
    If the key isn't found, the program exits.
    """

    keyMap = defineKeyboardMapping()
    categoryMap = defineKeyToCategory()
    img = cv2.imread(filename)
    print(img.shape)
    cv2.imshow(filename, img)

    while True:
        press = cv2.waitKey()
        key = keyMap.get(press, 27)
        if key == 27:
            exit(1)
        elif categoryMap.get(key, 0):
            category = categoryMap[key]
        else:
            print(f"{key} is not mapping to a category")
        cv2.destroyAllWindows()
        return category


def getListOfCategories(catMap: Dict[str, str] = defineKeyToCategory()):
    """
    Return a list of the possible call categories.
    """
    return list(catMap.values())


def createCategoryFoldersIfNonexistentIn(dir: str):
    """
    Create the category directories if they don't exist.
    """

    possibleCategories = getListOfCategories()

    for category in possibleCategories:
        if not os.path.exists(category):
            os.mkdir(category)


def findNonEmptyFolderIn(folder: str) -> str:
    """
    Return a subfolder within `folder` that contains pngs and is not
    one of the category folders we created.
    """
    categories = getListOfCategories(defineKeyToCategory())

    # get a list of folders
    folders = [x[0] for x in os.walk(folder)]
    folders_filtered = [folder for folder in folders if folder not in categories]

    # return a folder with at least one jpg
    for folder in folders_filtered:
        files = getListOfPNGsIn(folder)
        if files:
            return folder
    exit("No folders left with pngs")


def moveFileToCategory(file: str, category: str, topLevelDirectory: str):
    """
    Move `file` to the `category` directory in `topLevelDirectory`.
    """
    new_path = os.path.join(topLevelDirectory, category, os.path.basename(file))
    print(f"new path: {new_path}")
    print(f"old file: {file}")
    os.rename(file, new_path)


if __name__ == "__main__":

    topLevelDirectory = getFolder()
    createCategoryFoldersIfNonexistentIn(topLevelDirectory)
    directory = findNonEmptyFolderIn(topLevelDirectory)
    print(f"using {directory}")

    files = getListOfPNGsIn(directory)

    for file in files:
        print(f"Using file {file}")
        category = showImageToUser(os.path.join(directory, file))
        print(category)
        moveFileToCategory(os.path.join(directory, file), category, topLevelDirectory)
