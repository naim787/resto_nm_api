import"../chunks/CWj6FrbW.js";import"../chunks/69_IOA4Y.js";import{p as x,i as _,as as q,j,am as D,a as w,f as M,c as s,s as t,r as i,n as P,g,m as B,b as S}from"../chunks/COSkQsm4.js";import{e as E}from"../chunks/DN8omCT8.js";import{b as I}from"../chunks/fwLK5LFJ.js";import{i as N}from"../chunks/C6hNewRt.js";/* empty css                */import{s as T}from"../chunks/BCdqXXNJ.js";import{s as z,r as A}from"../chunks/BbDZAae-.js";import{I as C}from"../chunks/CrQ9F5MQ.js";function F(m,r){x(r,!0);/**
 * @license @lucide/svelte v0.525.0 - ISC
 *
 * ISC License
 *
 * Copyright (c) for portions of Lucide are held by Cole Bemis 2013-2022 as part of Feather (MIT). All other copyright (c) for Lucide are held by Lucide Contributors 2022.
 *
 * Permission to use, copy, modify, and/or distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 *
 */let u=A(r,["$$slots","$$events","$$legacy"]);const e=[["path",{d:"m12 19-7-7 7-7"}],["path",{d:"M19 12H5"}]];C(m,z({name:"arrow-left"},()=>u,{get iconNode(){return e},children:(d,o)=>{var a=q(),p=j(a);T(p,()=>r.children??D),w(d,a)},$$slots:{default:!0}})),_()}var G=M('<div class="w-[100vw] h-[100vh] flex flex-col justify-center items-center bg-gray-900"><a href="/menu_list"><!></a> <h1 class="text-3xl mb-5 text-white">Buat Menu Baru</h1> <form class="p-4 bg-gray-950 text-white w-full md:w-[50vw] h-[70vh] rounded-3xl flex flex-col justify-evenly items-center"><input name="name" type="text" class="w-full bg-gray-900 rounded-xl p-2 m-1 md:h-14" placeholder="nama menu" required/> <select name="category" class="w-full select bg-gray-900 md:h-14" required><option disabled selected>Category</option><option>Makanan</option><option>Minuman</option><option>Lainnya..</option></select> <input name="price" type="number" class="w-full bg-gray-900 rounded-xl p-2 m-1 md:h-14" placeholder="Price" required/> <input name="stock" type="number" class="w-full bg-gray-900 rounded-xl p-2 m-1 md:h-14" placeholder="Stock" required/> <input name="image_url" type="file" class="w-full bg-gray-900 rounded-xl p-2 m-1 md:h-14 file-input" required/> <div class="w-full"><label for="description"># Description</label></div> <textarea name="description" id="description" class="w-full h-52 bg-gray-900 rounded-xl p-2" required></textarea> <input type="hidden" name="admin_id" value="1233"/> <button type="submit" class="w-full h-14 bg-green-500 rounded-xl font-bold text-2xl">Buat</button></form></div>');function X(m,r){x(r,!1);const u="http://10.10.9.27:3001";let e=B();async function d(b){b.preventDefault();const k=new FormData(g(e));try{const n=await fetch(`${u}/create-products`,{method:"POST",body:k});if(!n.ok)throw new Error("Gagal kirim data ke server");const $=await n.json();alert("Produk berhasil dibuat!"),console.log($),g(e).reset()}catch(n){console.error(n),alert("Terjadi kesalahan saat submit.")}}N();var o=G(),a=s(o),p=s(a);F(p,{size:40,class:"p-2 rounded-md bg-gray-400 text-black absolute top-5 left-5"}),i(a);var l=t(a,4),c=t(s(l),2),f=s(c);f.value=f.__value="";var h=t(f);h.value=h.__value="makanan";var v=t(h);v.value=v.__value="minuman";var y=t(v);y.value=y.__value="lainnya",i(c),t(c,12),P(2),i(l),I(l,b=>S(e,b),()=>g(e)),i(o),E("submit",l,d),w(m,o),_()}export{X as component};
