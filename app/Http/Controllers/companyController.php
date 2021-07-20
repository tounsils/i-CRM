<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DataTables;
use App\Models\company;  // use user model
use Illuminate\Support\Facades\DB;

class companyController extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        $companies = company::all();
        return view('companyhome', compact('companies'));
        /*
        return view('companyhome', compact('companies'), [
            'compan' => DB::table('companies')->paginate(5)
        ]);
        */
        /*
        return view('user.index', [
            'users' => DB::table('users')->paginate(15)
        ]);
        */

/*
        if ($request->ajax()) {
            $data = company::latest()->get();
            return Datatables::of($data)
                    ->addIndexColumn()
                    ->addColumn('action', function($row){

                           $btn = '<a href="javascript:void(0)" data-toggle="tooltip"  data-id="'.$row->id.'" data-original-title="Edit" class="edit btn btn-primary btn-sm editcompany">Edit</a>';

                           $btn = $btn.' <a href="javascript:void(0)" data-toggle="tooltip"  data-id="'.$row->id.'" data-original-title="Delete" class="btn btn-danger btn-sm deletecompany">Delete</a>';

                            return $btn;
                    })
                    ->rawColumns(['action'])
                    ->make(true);
        }

        return view('companyhome');
        */
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
        company::updateOrCreate(['id' => $request->company_id],
                [
                    'company_name' => $request->company_name, 
                    'email' => $request->email,
                    'phone' => $request->phone
                ]);        

        return response()->json(['success'=>'company saved successfully!']);
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
        $company->company_name =  $request->get('company_name');
        $company->email = $request->get('email');
        $company->phone = $request->get('phone');
        $company->save();

        $companies = company::all();
        return view('companyhome', compact('companies'))->with('success', 'company updated!');
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
        //return response()->json(['success'=>'company deleted!']);
        $companies = company::all();
        return view('companyhome', compact('companies'));
    }


    public function getData()
    {
    $companies = company::paginate(5);
    return view('companyhome', compact('companies'));
    }
 
}

