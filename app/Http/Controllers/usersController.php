<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\user;  // use user model
use DataTables;
use Illuminate\Support\Facades\DB;
Use Redirect;
use Illuminate\Support\Facades\Auth;
use Session;

class usersController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $users = user::all();
        return view('users.index',compact('users'));
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        return view('createuser');
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {

        $user = new user;
        $user -> first_name = $request -> input('first_name');
        $user -> last_name = $request -> input('last_name');
        $user -> email = $request -> input('email');
        $user -> phone = $request -> input('phone');
        $user -> role_id = $request -> input('role_id');
        $user -> save();
        return redirect(route('usershome')) -> with('successMsg','User sucessfull added');

    }



    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        $user = user::find($id);
        return view('users.edit', compact('user'));        
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        $request->validate([
            'first_name'=>'required',
            'email'=>'required'
        ]);
        /*
            'first_name' => 'Admin',
            'last_name' => 'Admin',
            'email' => 'Admin@admin.com',
            'phone' => '0022115511',
            'role_id' => 1,
            'password' => bcrypt('root'),

        */

        $user = user::find($id);
        $user->first_name =  $request->first_name;
        $user->last_name =  $request->last_name;
        $user->email = $request->get('email');
        $user->phone = $request->get('phone');
        if (!empty($request->password)) {
            $user->password = bcrypt($request->password); //  update password only if modified
            //$user->password = $request->get('password'); //  update password only if modified
            }
        $user->save();
        
        $notification = array(
            'message' => "User (<b>" . $request->get('first_name') . "</b>) updated!",
            'alert-type' => 'success'
        );

        return redirect("users")->with($notification);    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        //
    }


    public function getData()
    {
    if(Auth::check()){
        $users = user::paginate(5);
        return view('users.index', compact('users'));
    }

    $notification = array(
        'message' => "You are not allowed to access users data!",
        'alert-type' => 'alert-success');

    return redirect("welcome")->with($notification);

}

}
