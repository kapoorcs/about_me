
/*
 * GET home page.
 */

exports.index = function(req, res){
  res.render('index', { title: 'Express' });
};

exports.helloworld = function(req, res){
    res.render("helloworld", {"title": "Yo Yo Hello !!"})
};

exports.homepage = function(req,res){
    res.render("helloworld", {"title": "Hello !!"})
};

// this method will get all data from userscollection 
exports.userlist = function(db){
    return function(req,res){
        var collection = db.get('userscollection');
        collection.find({},{},function(e, docs){
            res.render('userlist',{
                'userlist' : docs
            });
        });
        
    };
};

// this method will insert a new record in userscollection

exports.newuser = function(req, res){
    res.render('newuser', {'title': 'Add New User'})
};


exports.adduser= function(db){
    return function(req, res){
        // First get your form values 
        var userName = req.body.username;
        var userEmail = req.body.email;
        
        // second setup the collection
        var collection= db.get('userscollection');
        
        // submit to db 
        collection.insert({
            "username": userName,
            "email": userEmail
        }, function(err, doc){
            // error handling
            if(err) return res.send("There was problem adding the information");
            
            // if it worked then
            res.location("userlist");
            res.redirect("userlist");
            
            }
        );

    }
};