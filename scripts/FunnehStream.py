class Stream:
    def __init__(self, accName, source, streamKey):
        self.accName = accName
        self.source = source
        self.streamKey = streamKey
        
class SmartStream:
    def __init__(self, accName, sources, controller, streamKey):
        self.accName = accName
        self.sources = sources
        self.controller = controller
        self.streamKey = streamKey

def streamThread(tNumber, stream):
    global exitCalled, threadList
    while not exitCalled and threadList[tNumber] == True:
        None
        
    
def masterConsole():
    print("MP4-to-Twitch-Streamer by Sauftot.")
    print("Use help for a list of commands.")
    print("Warning!: Currently, only mp4 file loop streaming is supported.")
    
    global exitCalled
    
    while not exitCalled:
        global threadList
        streamList = []
        
        cmd = input()
        cmd = cmd.lower()
        params = cmd.split(' ')
        
        if params[0] == "new":
            None # add stream to list with source and streamkey
        elif params[0] == "load":
            None # load list of streams with source paths and streamkeys from file
        elif params[0] == "store":
            None # store list of streams with source paths and streamkeys to file
        elif params[0] == "unload":
            None # unload all streams from list loaded by file
        elif params[0] == "clear":
            None # clear list
        elif params[0] == "remove":
            None # remove specific stream from list
        elif params[0] == "list":
            print("List of streams contains:")
            for stream in streamList:
                print(stream)
        elif params[0] == "start":
            None # start specific stream or all
        elif params[0] == "stop":
            None # stop specific or all
        elif params[0] == "quit" or params[0] == "exit":
            exitCalled = True # stop all and exit
        elif params[0] == "help":
            None # show commands and context
          
          
exitCalled = False
threadList = []
masterConsole()