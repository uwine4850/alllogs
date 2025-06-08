package serverdto
import "github.com/uwine4850/foozy/pkg/router/rest"

import "github.com/uwine4850/foozy/pkg/router/form"

type mydto_BaseResponseMessage struct {
    rest.ImplementDTOMessage
    TypemydtoBaseResponseMessage any `dto:"TypemydtoBaseResponseMessage"`
    Ok bool `dto:"Ok"`
    Error string `dto:"Error"`
}

func NewMydtoBaseResponseMessage(
    error string,
    ok bool,
) *mydto_BaseResponseMessage {
    return &mydto_BaseResponseMessage{ 
	    Error: error,
	    Ok: ok,
    }
}

type mydto_RegisterMessage struct {
    rest.ImplementDTOMessage
    TypemydtoRegisterMessage any `dto:"TypemydtoRegisterMessage"`
    Username string `dto:"Username"`
    Password string `dto:"Password"`
    RepeatPassword string `dto:"RepeatPassword"`
}

func NewMydtoRegisterMessage(
    password string,
    repeatPassword string,
    username string,
) *mydto_RegisterMessage {
    return &mydto_RegisterMessage{ 
	    Password: password,
	    RepeatPassword: repeatPassword,
	    Username: username,
    }
}

type mydto_LoginMessage struct {
    rest.ImplementDTOMessage
    TypemydtoLoginMessage any `dto:"TypemydtoLoginMessage"`
    Username string `dto:"Username"`
    Password string `dto:"Password"`
}

func NewMydtoLoginMessage(
    password string,
    username string,
) *mydto_LoginMessage {
    return &mydto_LoginMessage{ 
	    Password: password,
	    Username: username,
    }
}

type mydto_LoginResponseMessage struct {
    rest.ImplementDTOMessage
    TypemydtoLoginResponseMessage any `dto:"TypemydtoLoginResponseMessage"`
    JWT string `dto:"JWT"`
    UID int `dto:"UID"`
    Error string `dto:"Error"`
}

func NewMydtoLoginResponseMessage(
    error string,
    jWT string,
    uID int,
) *mydto_LoginResponseMessage {
    return &mydto_LoginResponseMessage{ 
	    Error: error,
	    JWT: jWT,
	    UID: uID,
    }
}

type mydto_ProfileMessage struct {
    rest.ImplementDTOMessage
    TypemydtoProfileMessage any `dto:"TypemydtoProfileMessage"`
    User mydto_UserMessage `dto:"User"`
    Id int `dto:"Id"`
    UserId int `dto:"UserId"`
    Description string `dto:"Description"`
    Avatar string `dto:"Avatar"`
    Token string `dto:"Token"`
    Error string `dto:"Error"`
}

func NewMydtoProfileMessage(
    avatar string,
    description string,
    error string,
    id int,
    token string,
    user mydto_UserMessage,
    userId int,
) *mydto_ProfileMessage {
    return &mydto_ProfileMessage{ 
	    Avatar: avatar,
	    Description: description,
	    Error: error,
	    Id: id,
	    Token: token,
	    User: user,
	    UserId: userId,
    }
}

type mydto_UserMessage struct {
    rest.ImplementDTOMessage
    TypemydtoUserMessage any `dto:"TypemydtoUserMessage"`
    Id int `dto:"Id"`
    Username string `dto:"Username"`
}

func NewMydtoUserMessage(
    id int,
    username string,
) *mydto_UserMessage {
    return &mydto_UserMessage{ 
	    Id: id,
	    Username: username,
    }
}

type mydto_GenTokenMessage struct {
    rest.ImplementDTOMessage
    TypemydtoGenTokenMessage any `dto:"TypemydtoGenTokenMessage"`
    UserId int `dto:"UserId"`
}

func NewMydtoGenTokenMessage(
    userId int,
) *mydto_GenTokenMessage {
    return &mydto_GenTokenMessage{ 
	    UserId: userId,
    }
}

type mydto_TokenResponse struct {
    rest.ImplementDTOMessage
    TypemydtoTokenResponse any `dto:"TypemydtoTokenResponse"`
    Token string `dto:"Token"`
    Error string `dto:"Error"`
}

func NewMydtoTokenResponse(
    error string,
    token string,
) *mydto_TokenResponse {
    return &mydto_TokenResponse{ 
	    Error: error,
	    Token: token,
    }
}

type mydto_ProfileUpdateMessage struct {
    rest.ImplementDTOMessage
    TypemydtoProfileUpdateMessage any `dto:"TypemydtoProfileUpdateMessage"`
    PID int `dto:"PID"`
    Description string `dto:"Description"`
    Avatar form.FormFile `dto:"Avatar"`
    OldAvatarPath string `dto:"OldAvatarPath"`
    DelAvatar bool `dto:"DelAvatar"`
}

func NewMydtoProfileUpdateMessage(
    avatar form.FormFile,
    delAvatar bool,
    description string,
    oldAvatarPath string,
    pID int,
) *mydto_ProfileUpdateMessage {
    return &mydto_ProfileUpdateMessage{ 
	    Avatar: avatar,
	    DelAvatar: delAvatar,
	    Description: description,
	    OldAvatarPath: oldAvatarPath,
	    PID: pID,
    }
}

type mydto_NotificationInfoMessage struct {
    rest.ImplementDTOMessage
    TypemydtoNotificationInfoMessage any `dto:"TypemydtoNotificationInfoMessage"`
    PID string `dto:"PID"`
    Text string `dto:"Text"`
}

func NewMydtoNotificationInfoMessage(
    pID string,
    text string,
) *mydto_NotificationInfoMessage {
    return &mydto_NotificationInfoMessage{ 
	    PID: pID,
	    Text: text,
    }
}

type mydto_ProjectMessage struct {
    rest.ImplementDTOMessage
    TypemydtoProjectMessage any `dto:"TypemydtoProjectMessage"`
    Author mydto_ProjectAuthor `dto:"Author"`
    Id int `dto:"Id"`
    UserId int `dto:"UserId"`
    Name string `dto:"Name"`
    Description string `dto:"Description"`
    Error string `dto:"Error"`
}

func NewMydtoProjectMessage(
    author mydto_ProjectAuthor,
    description string,
    error string,
    id int,
    name string,
    userId int,
) *mydto_ProjectMessage {
    return &mydto_ProjectMessage{ 
	    Author: author,
	    Description: description,
	    Error: error,
	    Id: id,
	    Name: name,
	    UserId: userId,
    }
}

type mydto_ProjectAuthor struct {
    rest.ImplementDTOMessage
    TypemydtoProjectAuthor any `dto:"TypemydtoProjectAuthor"`
    PID int `dto:"PID"`
    Username string `dto:"Username"`
    Avatar string `dto:"Avatar"`
}

func NewMydtoProjectAuthor(
    avatar string,
    pID int,
    username string,
) *mydto_ProjectAuthor {
    return &mydto_ProjectAuthor{ 
	    Avatar: avatar,
	    PID: pID,
	    Username: username,
    }
}

type mydto_ProjectLogGroupMessage struct {
    rest.ImplementDTOMessage
    TypemydtoProjectLogGroupMessage any `dto:"TypemydtoProjectLogGroupMessage"`
    Id int `dto:"Id"`
    ProjectId int `dto:"ProjectId"`
    Name string `dto:"Name"`
    Description string `dto:"Description"`
    Error string `dto:"Error"`
}

func NewMydtoProjectLogGroupMessage(
    description string,
    error string,
    id int,
    name string,
    projectId int,
) *mydto_ProjectLogGroupMessage {
    return &mydto_ProjectLogGroupMessage{ 
	    Description: description,
	    Error: error,
	    Id: id,
	    Name: name,
	    ProjectId: projectId,
    }
}

