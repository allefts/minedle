import os

directory = '../recipes/'
 
# iterate over files in
# that directory
for filename in os.listdir(directory):
    f = directory  + filename
    os.system("mongoimport " + f + " -d craftable-items -c items --uri mongodb+srv://allefts:sa680108@cluster0.duhbvwj.mongodb.net/?retryWrites=true&w=majority")
