@extends('layouts.main')

@section('title', 'Users Home')

@section('content')
 
<main id="main">

    <!-- ======= company Section ======= -->
    <section id="user" class="user">

      <div class="container" data-aos="fade-up">

        <header class="section-header">
          <h2>Users</h2>
          <p>Listing</p>
        </header>

        <div class="container">
        
        @if (session('message'))
        <div class="alert session('alert-type')">
         {!! session('message') !!}
        </div>
        @endif

    <a class="btn btn-success" href="javascript:void(0)" id="createNewuser"> 
    <i class="ri-add-line"></i>Create New user (modal)</a>



<div >
    <a style="margin: 19px;" href="{{ route('users.create')}}" class="btn btn-primary">New user (blade)</a>
</div>  



<table class="table table-striped">
    <thead>
        <tr>
        <th>ID</th>
          <th>First</th>
          <th>Last</th>
          <th>Email</th>
          <th>Phone</th>
          <th colspan = 2>Actions</th>
        </tr>
    </thead>
    <tbody>
        @foreach($users as $user)
        <tr>
        <td>{{ $user->id }}</td>
          <td>{{ $user->first_name }}</td>
          <td>{{ $user->last_name }}</td>
          <td>{{ $user->email }}</td>
          <td>{{ $user->phone }}</td>
            <td>
                <a href="{{ route('users.edit',$user->id)}}" class="btn btn-edit">Edit</a>
            </td>
            <td>
                <form action="{{ route('users.destroy', $user->id)}}" method="post">
                  @csrf
                  @method('DELETE')
                  <button class="btn btn-danger" type="submit">Delete</button>
                </form>
            </td>
        </tr>
        @endforeach
    </tbody>
  </table>
<div>
</div>


<div class="col-sm-12">

  @if(session()->get('success'))
    <div class="alert alert-success">
      {{ session()->get('success') }}  
    </div>
  @endif
</div>

@endsection