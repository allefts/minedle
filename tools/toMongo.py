import os

directory = '../recipes/'
 
# iterate over files in
# that directory
for filename in os.listdir(directory):
    f = directory  + filename
    os.system("mongoimport " + f + " -d MinecraftItems -c Items --uri mongodb+srv://allefts:sa680108@cluster0.duhbvwj.mongodb.net/?retryWrites=true&w=majority")