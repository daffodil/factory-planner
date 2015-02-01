
package accounts


import (
	"strings"
	"errors"
	"github.com/jinzhu/gorm"
)

type Secretstruct struct {
	Secret string
}

// Authenticates and validate an User (currently only staff)
// error checking is done where as it can be called from many places
func AuthenticateUser(db gorm.DB, syslogin, secret string) (*ContactView, error, error) {

	ssyslogin := strings.TrimSpace( syslogin )
	if len(ssyslogin) <= 1 {
		return nil, errors.New("No `syslogin` or `syslogin` too short"), nil
	}

	ssecret := strings.TrimSpace( secret )
	if len(ssecret) <= 3 {
		return nil, errors.New( "No `password` or `password` to short"), nil
	}

	var user ContactView
	db.Table("v_contacts").Select(CONTACT_VIEW_COLS).Where("root = 1 and syslogin=?", syslogin).Scan(&user)

	if user.ContactId == 0 {
		return nil, errors.New("User not found"), nil
	}
	//if  user.CanLogin  {
	//	return nil, errors.New("User not allowed to login"), nil
	//}

	// recover passwords and compare.. TODO encrypt etc (but this is for legacy atmo)
	var stored_secret Secretstruct
	db.Table("contacts").Select("secret").Where("contact_id = ?", user.ContactId).Scan(&stored_secret)
	if stored_secret.Secret != secret {
		return nil, errors.New("Passwords not match"), nil
	}

	return &user, nil, nil
}
