import os
import tables
import opencv/core, opencv/highgui, opencv/imgproc

proc defineKeyboardMapping(): Table[int, string] =
    var intToKeyMapping = {
        48:  "0",
        49:  "1",
        50:  "2",
        51:  "3",
        52:  "4",
        53:  "5",
        54:  "6",
        55:  "7",
        56:  "8",
        57:  "9",
        119: "w",
        99:  "c",
        32:  " ",
        117: "u",
        }.toTable
    return intToKeyMapping

proc defineKeyToCategory(): Table[string, string] =
    var mapping = {
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
        }.toTable
    return mapping


proc getFolder(): string =
    var dir = getCurrentDir()
    return dir

proc getListOfFilesIn(folder: string): seq[string] =
    var files: seq[string] = @[]
    for file in walkFiles(folder & "/*.png"):
        files.add(file)
    return files

proc getListOfCategories(catMap: Table[string, string]): seq[string] =
    var categories: seq[string] = @[]
    for value in values(catMap):
        categories.add(value)
    return categories

proc createCategoryFoldersIfNonexistentIn(dir: string) =
    var possibleCategories = getListOfCategories(defineKeyToCategory())
    for category in possibleCategories:
        discard existsOrCreateDir(category)


proc findNonEmptyFolderIn(folder: string): string =

    var possibleFoldersToTest: seq[string] = @[]
    var categories = getListOfCategories(defineKeyToCategory())

    for kind, path in walkDir(folder):
        if not categories.contains(path.splitPath.tail):
            if existsDir(path.splitPath.tail):
                possibleFoldersToTest.add(path)
    
    for folder in possibleFoldersToTest:
        for file in walkFiles(folder & "/*.png"):
            return folder
    

proc isANonCategoricalFolder(folder: string): bool =
    var 
        isNotCat = true
        dirsWeCreated = getListOfCategories(defineKeyToCategory())
    for dir in dirsWeCreated:
        if dir == folder:
            isNotCat = false
    return isNotCat

proc showImageToUser(filename: string, delay = 1000) =
    var img = loadImage(fileName, 1)
    showImage(filename, img)
    discard waitKey(delay.cint)   # wait delay millisecs if key not pressed
    destroyAllWindows()

proc main() =
    var topLevelDirectory = getFolder()

    echo "topLevelDirectory is ", topLevelDirectory

    var keyboardMapping = defineKeyboardMapping()

    echo "\n\n"
    var categories = getListOfCategories(defineKeyToCategory())

    # create the folders if they don't exists
    createCategoryFoldersIfNonexistentIn(topLevelDirectory)

    # select a folder containing pngs
    var directory = findNonEmptyFolderIn(topLevelDirectory)
    echo "using ", directory

    var files = getListOfFilesIn(directory)

    for file in files:
        showImageToUser(file)

    

main()