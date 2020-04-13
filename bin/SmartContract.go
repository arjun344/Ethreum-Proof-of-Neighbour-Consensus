pragma solidity 0.4.21;
contract Linked {
    
    
    struct MedicalRecord {
        
        uint id;
        string disease;
        string reports;
        string doclicense;
        string verified;
        
    }
    
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
        uint MedicalRecordSize;
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
    
    mapping (uint => mapping (uint => MedicalRecord)) public MedOfUser;
    
    uint256 userCounter = 0;
    uint256 doctorCounter = 0;
    uint256 insurerCounter = 0;
    
    function strConcat(string _a, string _b, string _c, string _d, string _e) internal returns (string){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        bytes memory _bc = bytes(_c);
        bytes memory _bd = bytes(_d);
        bytes memory _be = bytes(_e);
        string memory abcde = new string(_ba.length + _bb.length + _bc.length + _bd.length + _be.length);
        bytes memory babcde = bytes(abcde);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) babcde[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) babcde[k++] = _bb[i];
        for (i = 0; i < _bc.length; i++) babcde[k++] = _bc[i];
        for (i = 0; i < _bd.length; i++) babcde[k++] = _bd[i];
        for (i = 0; i < _be.length; i++) babcde[k++] = _be[i];
        return string(babcde);
    }

    
    
    function AddMedicalRecord(uint _user, string _disease,string _report,string _doclicense) public {
        
        uint id = UserInfo[_user].MedicalRecordSize;
        MedicalRecord memory medical = MedicalRecord(id,_disease,_report,_doclicense,"0");
        MedOfUser[_user][UserInfo[_user].MedicalRecordSize] = medical;
        UserInfo[_user].MedicalRecordSize++;
        id++;
    }
    
    function UpdateReport(uint _userid,uint _recordid,string _rep) public
    {
       
        string storage _report =  MedOfUser[_userid][_recordid].reports;
        string memory k = strConcat(_report,_rep,"","","");
        MedOfUser[_userid][_recordid].reports = k;
        
    }
    
    function UpdateReportStats(uint _userid,uint _recordid,string _stats) public
    {
        MedOfUser[_userid][_recordid].verified = _stats;
    }
    
    // Sets the profile of a user 
    function setUserProfile(string _name, string  _email,string _uniqueid,string _password) public {
        
        User memory user = User(_name, _email,_uniqueid,0,0,0);
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

