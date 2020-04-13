pragma solidity 0.4.21;
contract Linked {
    
    //Insurer profile
    struct Insurer {
        
        string name;
        string email;
        string license;
        
        
    }
     struct InsurerLogin{
        
        string email;
        string password;
        uint UsrOfInsSize;
        
    }
    
    // Doctor profile
    struct Doctor {
        
        string name;
        string email;
        string uniqueid;
        string license;
        string organization;
        
    }
    struct DoctorLogin {
        
        string email;
        string password;
        uint PatOfDocSize;
    }
    
    // User profile
    struct User {
        string name;
        string email;
        string uniqueid;
        uint DoccOfUserSize;
        uint InsOfUserSize;
    }
    struct UserLogin {
        string email;
        string password;
    }
    
    mapping (uint => mapping (uint => Doctor)) public DoccOfUser;
    mapping (uint => mapping (uint => Insurer)) public InsOfUser;
    
    mapping (uint => mapping (uint => User)) public PatOfDoc;
    mapping (uint => mapping (uint => User)) public UsrOfIns;
    
    // Each address is linked to a user with name, occupation and bio
    mapping(uint256 => User) public UserInfo;
    mapping(uint256 => UserLogin) public UserLoginInfo;
    mapping(uint256 => Doctor) public DoctorInfo;
    mapping(uint256 => DoctorLogin) public DoctorLoginInfo;
    mapping(uint256 => Insurer) public InsurerInfo;
    mapping(uint256 => InsurerLogin) public InsurerLoginInfo;
    
    uint256 userCounter = 0;
    uint256 doctorCounter = 0;
    uint256 insurerCounter = 0;
    
    // Sets the profile of a user 
    function setUserProfile(string _name, string  _email,string _uniqueid,string _password) public {
        
        User memory user = User(_name, _email,_uniqueid,0,0);
        UserInfo[userCounter] = user;
        
        UserLogin memory ulog = UserLogin(_email,_password);
        UserLoginInfo[userCounter] = ulog;
        
        userCounter++;
    }
    
    function AddDocToPat(uint _pat,uint _doc) public {
        
        DoccOfUser[_pat][UserInfo[_pat].DoccOfUserSize] = DoctorInfo[_doc];
        UserInfo[_pat].DoccOfUserSize++;
        
    }
    
     function AddPatToDoc(uint _doc,uint _pat) public {
        
        PatOfDoc[_doc][DoctorLoginInfo[_doc].PatOfDocSize] = UserInfo[_pat];
        DoctorLoginInfo[_doc].PatOfDocSize++;
        
    }
    
    function AddInsToUsr(uint _usr,uint _ins) public {
        
        InsOfUser[_usr][UserInfo[_usr].InsOfUserSize] = InsurerInfo[_ins];
        UserInfo[_usr].InsOfUserSize++;
        
    }
    
    function AddUsrToIns(uint _ins,uint _usr) public {
        
        UsrOfIns[_ins][InsurerLoginInfo[_ins].UsrOfInsSize] = UserInfo[_usr];
        InsurerLoginInfo[_ins].UsrOfInsSize++;
        
    }
    
    
    function setDoctorProfile(string _name, string  _email,string _uniqueid,string _license,string _organisation,string _password) public {
        
        Doctor memory doctor = Doctor(_name, _email,_uniqueid,_license,_organisation);
        DoctorInfo[doctorCounter] = doctor;
        
        
        DoctorLogin memory dlog = DoctorLogin(_email,_password,0);
        DoctorLoginInfo[doctorCounter] = dlog;
        
        doctorCounter++;
    }
    
    
    
    
    function setInsurerProfile(string _name, string  _email,string _license,string _password) public {
        
        Insurer memory insurer = Insurer(_name, _email,_license);
        InsurerInfo[insurerCounter] = insurer;
        
        InsurerLogin memory ilog = InsurerLogin(_email,_password,0);
        InsurerLoginInfo[insurerCounter] = ilog;
        
        insurerCounter++;
    }
    
    
    function getUserCounter() view public returns (uint256) {
        
        return userCounter;
    }
    
    function getDoctorCounter() view public returns (uint256) {
        
        return doctorCounter;
    }
    
    function getInsurerCounter() view public returns (uint256) {
        
        return insurerCounter;
    }
    
    
   

}

