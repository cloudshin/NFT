var express = require('express');
var router = express.Router();
const cc = require('../public/js/cc');

const USER_COOKIE_KEY = 'USER';

router.post('/', async(req, res, next) => {
    console.log("/in transferFrom")
    const userCookie  = req.cookies[USER_COOKIE_KEY];
    const userData = JSON.parse(userCookie);

    const id = userData.username    
    const fungusid = req.body.fungusid
    const to_id =  req.body.to_id
    //from_id " Carolina,C=US"까지 읽어올 수 있도록 수정
    const from_id = req.body.from_id
    const from_id2 = from_id + " Carolina,C=US"
    var args = [from_id2, to_id, fungusid];
    console.log(`to_id: ${to_id}`);
    console.log(`from_id: ${from_id2}`);
    var result = await cc.cc_call(id, "TransferFrom", args)
    console.log(result)
    res.redirect('/');
});

module.exports = router;