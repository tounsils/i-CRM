<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DataTables;
use App\Models\company;  // use user model
use Illuminate\Support\Facades\DB;
Use Redirect;
use Illuminate\Support\Facades\Auth;
use Session;

class companyController extends Controller
{

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        // See getData() function
        if(Auth::check()){
            $companies = company::all();
            return view('company.index', compact('companies'));
        }
  
        return redirect("welcome")->withSuccess('You are not allowed to access');
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        $company = new company();
        $company->company_name = $request->company_name; 
        $company->email = $request->email;
        $company->phone = $request->phone;

    /*
        $company = company::updateOrCreate(['id' => $request->company_id],
                [
                    'company_name' => $request->company_name, 
                    'email' => $request->email,
                    'phone' => $request->phone
                ]); 
                */
                $company->save();       
                
                $companies = company::all();
                return view('company.index', compact('companies'))->with('success', 'company updated!');

        //return response()->json(['success'=>'company saved successfully!']);
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        $companies = company::paginate(5);
        return view('company.index', compact('companies'));
/*
        $company = company::find($id);
        return view('company.index', compact('company'));
        */
        }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        $company = company::find($id);
        return view('company.edit', compact('company'));        
        //$company = company::find($id);
        //return response()->json($company);
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
            'company_name'=>'required',
            'email'=>'required'
        ]);

        $company = company::find($id);
        $company->company_name =  $request->company_name;
        $company->email = $request->get('email');
        //$employee->name = $request->fullname;

        $company->phone = $request->get('phone');
        $company->save();
        
        $notification = array(
            'message' => "Company (<b>" . $request->get('company_name') . "</b>) updated!",
            'alert-type' => 'success'
        );

        return redirect("company")->with($notification);
        //companyController::getData();
    }
/*
            $table->string('company_name');
            $table->string('email');
            $table->string('phone');
*/

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        company::find($id)->delete();
        $notification = array(
            'message' => "Company deleted!",
            'alert-type' => 'success'
        );
        
        return Redirect::back()->with($notification);

    }


    public function getData()
    {
    if(Auth::check()){
        $companies = company::paginate(5);
        return view('company.index', compact('companies'));
    }

    $notification = array(
        'message' => "You are not allowed to access Company data!",
        'alert-type' => 'alert-success');

    return redirect("welcome")->with($notification);

}
 
}

